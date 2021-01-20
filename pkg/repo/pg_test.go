package repo_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/potalestor/custom-wallet/pkg/cfg"
	"github.com/potalestor/custom-wallet/pkg/db"
	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
	"github.com/stretchr/testify/assert"
)

var config = &cfg.Config{
	Database: cfg.Database{
		User:     "postgres",
		Password: "postgres",
		DB:       "wallet_test",
		Host:     "localhost",
		Port:     5432,
	},

	Migration: cfg.Migration{
		Enabled: true,
		Path:    "../../scripts",
	},
}

func TestMain(m *testing.M) {
	migrationdb := db.NewPostgresDB(config)
	if err := migrationdb.Open(); err != nil {
		log.Fatalf("migration does not initialize: %v\n%+v", err, config)
	}

	if err := migrationdb.Migrate(); err != nil {
		log.Fatalf("migration does not perform: %v\n%+v", err, config)
	}

	migrationdb.Close()

	os.Exit(m.Run())
}

func TestPgStorage_Deposit(t *testing.T) {
	ctx := context.Background()

	storage := repo.NewPgStorage(config)
	if err := storage.Open(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		assert.NoError(t, storage.Clear(ctx))
		assert.NoError(t, storage.Close())
	}()

	tests := []struct {
		name    string
		wallet  *model.Wallet
		amount  model.USD
		wantErr bool
	}{
		{"correct", &model.Wallet{Name: "wd1"}, 10, false},
		{"zero", &model.Wallet{Name: "wd2"}, 0, true},
		{"negative", &model.Wallet{Name: "wd3"}, -10, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, storage.CreateWallet(ctx, tt.wallet))

			if tt.wantErr {
				assert.Error(t, storage.Deposit(ctx, tt.wallet, tt.amount))

				return
			}

			assert.NoError(t, storage.Deposit(ctx, tt.wallet, tt.amount))
			assert.NotEqual(t, 0, tt.wallet)
			assert.Equal(t, tt.amount, tt.wallet.Account)
		})
	}
}

func TestPgStorage_Transfer(t *testing.T) {
	ctx := context.Background()

	storage := repo.NewPgStorage(config)
	if err := storage.Open(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		assert.NoError(t, storage.Clear(ctx))
		assert.NoError(t, storage.Close())
	}()

	tests := []struct {
		name    string
		src     *model.Wallet
		dst     *model.Wallet
		amount  model.USD
		wantErr bool
	}{
		{
			"correct",
			&model.Wallet{Name: "w1", Account: 10},
			&model.Wallet{Name: "w2", Account: 20},
			10,
			false,
		},
		{
			"negative",
			&model.Wallet{Name: "w3", Account: 20},
			&model.Wallet{Name: "w4", Account: 20},
			30,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, storage.CreateWallet(ctx, tt.src))
			assert.NoError(t, storage.Deposit(ctx, tt.src, tt.src.Account))
			assert.NoError(t, storage.CreateWallet(ctx, tt.dst))
			assert.NoError(t, storage.Deposit(ctx, tt.dst, tt.dst.Account))

			if tt.wantErr {
				assert.Error(t, storage.Transfer(ctx, tt.src, tt.dst, tt.amount))

				return
			}

			assert.NoError(t, storage.Transfer(ctx, tt.src, tt.dst, tt.amount))
		})
	}
}

func TestPgStorage_CreateWallet(t *testing.T) {
	ctx := context.Background()

	storage := repo.NewPgStorage(config)
	if err := storage.Open(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		assert.NoError(t, storage.Clear(ctx))
		assert.NoError(t, storage.Close())
	}()

	tests := []struct {
		name    string
		wallet  *model.Wallet
		wantErr bool
	}{
		{"correct", &model.Wallet{Name: "wc1"}, false},
		{"zero", &model.Wallet{Name: "wc1"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Error(t, storage.CreateWallet(ctx, tt.wallet))

				return
			}

			assert.NoError(t, storage.CreateWallet(ctx, tt.wallet))
		})
	}
}

func TestPgStorage_Report(t *testing.T) {
	ctx := context.Background()

	storage := repo.NewPgStorage(config)
	if err := storage.Open(); err != nil {
		t.Fatal(err)
	}

	defer func() {
		assert.NoError(t, storage.Clear(ctx))
		assert.NoError(t, storage.Close())
	}()

	f := model.NewFilter()
	f.WalletName = "w1"

	_, err := storage.Report(ctx, f)

	assert.Error(t, err)
}
