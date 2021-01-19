package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase_Validate(t *testing.T) {
	tests := []struct {
		name    string
		db      Database
		wantErr bool
	}{
		{"correct", Database{"postgres", "postgres", "wallet", "localhost", 5432}, false},
		{"correct", Database{"postgres", "postgres", "", "localhost", 5432}, false},
		{"invalid", Database{"", "", "", "", 0}, true},
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
	d := Database{"postgres", "postgres", "wallet", "localhost", 5432}
	assert.Equal(t, expected, d.String())

	expected = `host=localhost port=5432 user=postgres password=postgres sslmode=disable`
	d = Database{"postgres", "postgres", "", "localhost", 5432}
	assert.Equal(t, expected, d.String())
}

func TestMigration_Validate(t *testing.T) {

	tests := []struct {
		name    string
		m       Migration
		wantErr bool
	}{
		{"correct", Migration{true, "../../migration"}, false},
		{"correct", Migration{false, ""}, false},
		{"invalid", Migration{true, "../../mgrtn/"}, true},
		{"invalid", Migration{true, ""}, true},
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
		config  Config
		wantErr bool
	}{
		{"correct",
			Config{
				Database{"postgres", "postgres", "wallet", "localhost", 5432},
				Migration{true, "../../migration"},
			},
			false},

		{"invalid",
			Config{
				Database{"", "", "", "", 0},
				Migration{true, "../../migration"},
			},
			true},

		{"invalid",
			Config{
				Database{"postgres", "postgres", "wallet", "localhost", 5432},
				Migration{true, "../../mgrtn/"},
			},
			true},
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
	d := Database{"postgres", "postgres", "wallet", "localhost", 5432}
	assert.Equal(t, expected, d.URI())
}
