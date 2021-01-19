package api

import (
	"github.com/potalestor/custom-wallet/pkg/app"
)

type Transfer struct {
	wallet *app.Wallet
}

func NewTransfer(wallet *app.Wallet) *Transfer {
	return &Transfer{wallet: wallet}
}
