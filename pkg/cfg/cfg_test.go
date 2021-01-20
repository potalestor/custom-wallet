package cfg_test

import (
	"testing"

	"github.com/potalestor/custom-wallet/pkg/cfg"
	"github.com/stretchr/testify/assert"
)

func TestDatabase_Validate(t *testing.T) {
	tests := []struct {
		name    string
		db      cfg.Database
		wantErr bool
	}{
		{"correct", cfg.Database{"postgres", "postgres", "wallet", "localhost", 5432}, false},
		{"correct", cfg.Database{"postgres", "postgres", "", "localhost", 5432}, false},
		{"invalid", cfg.Database{"", "", "", "", 0}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Error(t, tt.db.Validate())

				return
			}

			assert.NoError(t, tt.db.Validate())
		})
	}
}

func TestDatabase_String(t *testing.T) {
	expected := `host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable`
	d := cfg.Database{"postgres", "postgres", "wallet", "localhost", 5432}
	assert.Equal(t, expected, d.String())

	expected = `host=localhost port=5432 user=postgres password=postgres sslmode=disable`
	d = cfg.Database{"postgres", "postgres", "", "localhost", 5432}
	assert.Equal(t, expected, d.String())
}

func TestMigration_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       cfg.Migration
		wantErr bool
	}{
		{"correct", cfg.Migration{true, "../../scripts"}, false},
		{"correct", cfg.Migration{false, ""}, false},
		{"invalid", cfg.Migration{true, "../../mgrtn/"}, true},
		{"invalid", cfg.Migration{true, ""}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Error(t, tt.m.Validate())

				return
			}

			assert.NoError(t, tt.m.Validate())
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  cfg.Config
		wantErr bool
	}{
		{
			"correct",
			cfg.Config{
				cfg.Database{"postgres", "postgres", "wallet", "localhost", 5432},
				cfg.Migration{true, "../../scripts"},
				cfg.Logger{},
				cfg.Web{},
			},
			false,
		},

		{
			"invalid",
			cfg.Config{
				cfg.Database{"", "", "", "", 0},
				cfg.Migration{true, "../../scripts"},
				cfg.Logger{},
				cfg.Web{},
			},
			true,
		},

		{
			"invalid",
			cfg.Config{
				cfg.Database{"postgres", "postgres", "wallet", "localhost", 5432},
				cfg.Migration{true, "../../mgrtn/"},
				cfg.Logger{},
				cfg.Web{},
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Error(t, tt.config.Validate())

				return
			}

			assert.NoError(t, tt.config.Validate())
		})
	}
}

func TestDatabase_URI(t *testing.T) {
	expected := `postgres://postgres:postgres@localhost:5432/wallet?sslmode=disable`
	d := cfg.Database{"postgres", "postgres", "wallet", "localhost", 5432}
	assert.Equal(t, expected, d.URI())
}
