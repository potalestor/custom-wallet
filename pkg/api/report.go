package api

import (
	"github.com/potalestor/custom-wallet/pkg/app"
)

type Report struct {
	wallet *app.Wallet
}

func NewReport(wallet *app.Wallet) *Report {
	return &Report{wallet: wallet}
}
