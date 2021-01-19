package app

import (
	"github.com/potalestor/custom-wallet/pkg/repo"
	"github.com/potalestor/custom-wallet/pkg/service"
)

// Wallet is main app.
type Wallet struct {
	Wallet   *service.Wallet
	Transfer *service.Transfer
	Report   *service.Report
}

// NewWallet returns new app.
func NewWallet(repository repo.Repository) *Wallet {
	return &Wallet{
		Wallet:   service.NewWallet(repository),
		Transfer: service.NewTransfer(repository),
		Report:   service.NewReport(repository),
	}
}
