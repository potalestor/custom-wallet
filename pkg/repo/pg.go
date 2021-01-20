package repo

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/potalestor/custom-wallet/pkg/cfg"
	"github.com/potalestor/custom-wallet/pkg/model"
)

const (
	sqlDriver          = `postgres`
	sqlCreateWallet    = `INSERT INTO wallets (name) VALUES ($1) RETURNING id;`
	sqlChangeAccount   = `UPDATE wallets SET account = account + ($1) WHERE id = ($2) RETURNING account;`
	sqlGetWalletByName = `SELECT id, account FROM wallets WHERE name = ($1);`
	sqlGetWalletByID   = `SELECT name, account FROM wallets WHERE id = ($1);`
	sqlCreateOperation = `INSERT INTO transactions (wallet_id, operation, amount) VALUES ($1, $2, $3) RETURNING id;`
	sqlTruncate        = `TRUNCATE transactions, wallets RESTART IDENTITY;`
	sqlReport          = `SELECT operation, created, amount FROM public.transactions where wallet_id=($1) and operation & ($2) > 0 and created between ($3) and ($4);`
	isolationLevel     = sql.LevelRepeatableRead
)

var (
	ErrWalletExists       = errors.New(`wallet already exists`)
	ErrWalletNotExist     = errors.New(`wallet does not exist`)
	ErrZeroAccount        = errors.New(`the amount on the account can't be less than zero`)
	ErrInvalidOperation   = errors.New(`invalid operation`)
	ErrInvalidTransaction = errors.New(`invalid transaction`)
)

// PgStorage implements repository interface.
type PgStorage struct {
	config *cfg.Config
	db     *sql.DB
}

// NewPgStorage creates new instance.
func NewPgStorage(config *cfg.Config) *PgStorage {
	return &PgStorage{config: config}
}

// Open database.
func (r *PgStorage) Open() error {
	db, err := sql.Open(sqlDriver, r.config.Database.String())
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	r.db = db

	return nil
}

// Close database.
func (r *PgStorage) Close() error {
	if r.db != nil {
		return r.db.Close()
	}

	return nil
}

// Clear database.
func (r *PgStorage) Clear(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, sqlTruncate)

	return err
}

// CreateWallet creates new wallet and get ID from database or error if does not exist.
func (r *PgStorage) CreateWallet(ctx context.Context, wallet *model.Wallet) error {
	err := r.db.QueryRowContext(ctx, sqlCreateWallet, wallet.Name).Scan(&wallet.ID)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == pq.ErrorCode("23505") {
			return errors.Wrap(ErrWalletExists, wallet.String())
		}
	}

	return err
}

// GetWalletByID return wallet by ID or error if does not exist.
func (r *PgStorage) GetWalletByID(ctx context.Context, wallet *model.Wallet) error {
	err := r.db.QueryRowContext(ctx, sqlGetWalletByID, wallet.ID).Scan(&wallet.Name, &wallet.Account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(ErrWalletNotExist, wallet.String())
		}
	}

	return err
}

// GetWalletByName return wallet by Name or error.
func (r *PgStorage) GetWalletByName(ctx context.Context, wallet *model.Wallet) error {
	err := r.db.QueryRowContext(ctx, sqlGetWalletByName, wallet.Name).Scan(&wallet.ID, &wallet.Account)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(ErrWalletNotExist, wallet.String())
		}
	}

	return err
}

// Transfer perform from src wallet to dst wallet.
func (r *PgStorage) Transfer(ctx context.Context, src, dst *model.Wallet, amount model.USD) error {
	op1 := &model.Transaction{
		Wallet:    src.ID,
		Operation: model.Withdraw,
		Amount:    amount,
	}

	op2 := &model.Transaction{
		Wallet:    dst.ID,
		Operation: model.Deposit,
		Amount:    amount,
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.changeAccount(ctx, tx, src, -amount); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.createOperation(ctx, tx, op1); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.changeAccount(ctx, tx, dst, amount); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.createOperation(ctx, tx, op2); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	return tx.Commit()
}

// Deposit wallet.
func (r *PgStorage) Deposit(ctx context.Context, wallet *model.Wallet, amount model.USD) error {
	op := &model.Transaction{
		Wallet:    wallet.ID,
		Operation: model.Deposit,
		Amount:    amount,
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: isolationLevel})
	if err != nil {
		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.changeAccount(ctx, tx, wallet, amount); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	if err := r.createOperation(ctx, tx, op); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}

		return errors.Wrap(ErrInvalidTransaction, err.Error())
	}

	return tx.Commit()
}

// Report wallet.
func (r *PgStorage) Report(ctx context.Context, filter *model.Filter) (model.Reports, error) {
	wallet := &model.Wallet{Name: filter.WalletName}

	if err := r.GetWalletByName(ctx, wallet); err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlReport,
		wallet.ID,
		filter.Operation,
		filter.DateRange[0], filter.DateRange[1],
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports model.Reports
	for rows.Next() {
		var report model.Report
		if err := rows.Scan(
			&report.Operation,
			&report.Created,
			&report.Amount,
		); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}

	rerr := rows.Close()
	if rerr != nil {
		return reports, err
	}

	return reports, rows.Err()
}

func (r *PgStorage) changeAccount(ctx context.Context, tx *sql.Tx, wallet *model.Wallet, amount model.USD) error {
	err := tx.QueryRowContext(ctx, sqlChangeAccount, amount.Float64(), wallet.ID).
		Scan(&wallet.Account)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == pq.ErrorCode("23514") {
			return errors.Wrap(ErrZeroAccount, wallet.String())
		}

		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(ErrWalletNotExist, wallet.String())
		}
	}

	return err
}

func (r *PgStorage) createOperation(ctx context.Context, tx *sql.Tx, transaction *model.Transaction) error {
	err := tx.QueryRowContext(ctx,
		sqlCreateOperation,
		transaction.Wallet,
		transaction.Operation,
		transaction.Amount.Float64()).
		Scan(&transaction.ID)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == pq.ErrorCode("23503") {
			return errors.Wrap(ErrInvalidOperation, pgErr.Detail)
		}
	}

	return err
}
