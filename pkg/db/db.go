package db

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/potalestor/custom-wallet/pkg/cfg"
)

const (
	sqlDriver = `postgres`

	sqlCreateTestDB = `CREATE DATABASE `
)

var ErrDatabaseNotExist = errors.New(`database does not exist`)

// PostgresDB ...
type PostgresDB struct {
	config *cfg.Config
	db     *sql.DB
}

// NewPostgresDB creates new instance.
func NewPostgresDB(config *cfg.Config) *PostgresDB {
	return &PostgresDB{config: config}
}

// Open database.
func (p *PostgresDB) Open() error {
	db, err := sql.Open(sqlDriver, p.config.Database.String())
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == pq.ErrorCode("3D000") {
			if e := p.createDB(); e != nil {
				return errors.Wrap(ErrDatabaseNotExist, pgErr.Message)
			}

			return p.Open()
		}

		return err
	}

	p.db = db

	return nil
}

// Close database.
func (p *PostgresDB) Close() error {
	if p.db != nil {
		return p.db.Close()
	}

	return nil
}

func (p *PostgresDB) Migrate() error {
	m := NewMigration(p.config)

	return m.Up()
}

func (p *PostgresDB) createDB() error {
	name := p.config.Database.DB
	p.config.Database.DB = ""

	temp := NewPostgresDB(p.config)
	if err := temp.Open(); err != nil {
		return err
	}

	defer temp.Close()

	if _, err := temp.db.Exec(sqlCreateTestDB + name); err != nil {
		return err
	}

	p.config.Database.DB = name

	return nil
}
