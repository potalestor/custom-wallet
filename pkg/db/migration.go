package db

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // for migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/potalestor/custom-wallet/pkg/cfg"
)

const formatFile = `file://%s`

// Migration implements golang-migrate for Postgresql.
type Migration struct {
	config *cfg.Config
}

// NewMigration returns new instance.
func NewMigration(config *cfg.Config) *Migration {
	return &Migration{config: config}
}

// Up migrate.
func (m *Migration) Up() error {
	if !m.config.Migration.Enabled {
		return nil
	}

	mg, err := migrate.New(
		fmt.Sprintf(formatFile, m.config.Migration.Path),
		m.config.Database.URI())
	if err != nil {
		return err
	}

	err = mg.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

// Down migrate.
func (m *Migration) Down() error {
	if !m.config.Migration.Enabled {
		return nil
	}

	mg, err := migrate.New(
		fmt.Sprintf(formatFile, m.config.Migration.Path),
		m.config.Database.URI())
	if err != nil {
		return err
	}

	err = mg.Down()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
