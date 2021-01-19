package repo_test

import (
	"testing"

	"github.com/potalestor/custom-wallet/pkg/repo"
	"github.com/stretchr/testify/assert"
)

func TestMigration_Up(t *testing.T) {
	assert.NoError(t, config.Validate())

	m := repo.NewMigration(config)
	assert.NoError(t, m.Up())
}

// func TestMigration_Down(t *testing.T) {
// 	assert.NoError(t, config.Validate())

// 	m := repo.NewMigration(config)
// 	assert.NoError(t, m.Down())
// }
