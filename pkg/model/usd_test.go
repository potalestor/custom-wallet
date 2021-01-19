package model_test

import (
	"testing"

	"github.com/potalestor/custom-wallet/pkg/model"
)

func TestUSD_Parse(t *testing.T) {
	tests := []struct {
		usd     string
		wantErr bool
	}{
		{`12.34`, false},
		{`0.34`, false},
		{`0.0`, false},
		{`0.00`, false},
		{`123`, false},
		{`12345`, false},
		{`1,345.00`, false},
		{`$123,456,789.99`, false},
		{`1,345`, false},
		{`1,2345`, true},
		{`12.34$`, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.usd, func(t *testing.T) {
			t.Parallel()
			var usd model.USD
			if err := usd.Parse(tt.usd); (err != nil) != tt.wantErr {
				t.Errorf("USD.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUSD_String(t *testing.T) {
	tests := []struct {
		name string
		usd  model.USD
	}{
		{"$12.34", model.USD(12.34)},
		{"$1234.56", model.USD(1234.56)},
		{"$12.00", model.USD(12)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.usd.String(); got != tt.name {
				t.Errorf("USD.String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestErrInvalidUSDFormat_Error(t *testing.T) {
	tests := []struct {
		name   string
		format string
	}{
		{"invalid USD format: 12.34$. Required: $12.34 or $1,234.56", "12.34$"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := model.NewErrInvalidUSDFormat(tt.format)
			if got := e.Error(); got != tt.name {
				t.Errorf("ErrInvalidUSDFormat.Error() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestUSD_Float64(t *testing.T) {
	tests := []struct {
		name string
		u    model.USD
		want float64
	}{
		{"$12.34", model.USD(12.34), 12.34},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.u.Float64(); got != tt.want {
				t.Errorf("USD.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
