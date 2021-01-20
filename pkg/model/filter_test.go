package model_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFilter_Validate(t *testing.T) {
	f := model.NewFilter()
	f.WalletName = "w1"

	assert.NoError(t, f.Validate())

	f = model.NewFilter()
	f.WalletName = ""

	assert.Error(t, f.Validate())

	f = model.NewFilter()
	f.WalletName = "w1"
	f.Operation = model.Unsupported

	assert.Error(t, f.Validate())

	f = model.NewFilter()
	f.WalletName = "w1"
	f.DateRange[1] = time.Time{}

	assert.Error(t, f.Validate())

	f = model.NewFilter()
	f.WalletName = "w1"
	f.DateRange[0] = time.Now().Add(10 * 24 * time.Hour)

	assert.Error(t, f.Validate())
}

func TestFilter_String(t *testing.T) {
	f := model.NewFilter()
	f.WalletName = "w1"

	assert.NoError(t, f.Validate())
	fmt.Println(f)
}
