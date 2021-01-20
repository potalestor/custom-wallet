package db_test

import (
	"testing"

	"github.com/potalestor/custom-wallet/pkg/cfg"
	"github.com/potalestor/custom-wallet/pkg/db"
	"github.com/stretchr/testify/assert"
)

var config = &cfg.Config{
	Database: cfg.Database{
		User:     "postgres",
		Password: "postgres",
		DB:       "migration_test",
		Host:     "localhost",
		Port:     5432,
	},

	Migration: cfg.Migration{
		Enabled: true,
		Path:    "../../scripts",
	},
}

func TestMigration_Up(t *testing.T) {
	assert.NoError(t, config.Validate())

	migrationdb := db.NewPostgresDB(config)
	assert.NoError(t, migrationdb.Open())

	m := db.NewMigration(config)
	assert.NoError(t, m.Up())

	assert.NoError(t, migrationdb.Close())

}

func TestMigration_Down(t *testing.T) {
	assert.NoError(t, config.Validate())

	migrationdb := db.NewPostgresDB(config)
	assert.NoError(t, migrationdb.Open())

	m := db.NewMigration(config)
	assert.NoError(t, m.Down())

	assert.NoError(t, migrationdb.Close())
}
