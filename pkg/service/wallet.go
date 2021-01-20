package service

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
)

// Wallet service.
type Wallet struct {
	repository repo.Repository
}

// NewWallet returns new instance.
func NewWallet(repository repo.Repository) *Wallet {
	return &Wallet{repository: repository}
}

// CreateWallet use case.
func (w *Wallet) CreateWallet(name string) (*model.Wallet, error) {
	wallet := model.Wallet{Name: name}

	if err := w.repository.CreateWallet(context.Background(), &wallet); err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Deposit use case.
func (w *Wallet) Deposit(name string, amount model.USD) (*model.Wallet, error) {
	if err := amount.Validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()

	wallet := model.Wallet{Name: name}

	if err := w.repository.GetWalletByName(ctx, &wallet); err != nil {
		return nil, err
	}

	if err := w.repository.Deposit(context.Background(), &wallet, amount); err != nil {
		return nil, err
	}

	return &wallet, nil
}
