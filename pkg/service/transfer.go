package service

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
)

// Transfer service.
type Transfer struct {
	repository repo.Repository
}

// NewTransfer returns new instance.
func NewTransfer(repository repo.Repository) *Transfer {
	return &Transfer{repository: repository}
}

// Transfer use case.
func (t *Transfer) Transfer(src, dst string, amount model.USD) ([]*model.Wallet, error) {
	if err := amount.Validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()

	srcWallet := model.Wallet{Name: src}

	if err := t.repository.GetWalletByName(ctx, &srcWallet); err != nil {
		return nil, err
	}

	dstWallet := model.Wallet{Name: dst}

	if err := t.repository.GetWalletByName(ctx, &dstWallet); err != nil {
		return nil, err
	}

	return []*model.Wallet{
		&srcWallet,
		&dstWallet,
	}, t.repository.Transfer(ctx, &srcWallet, &dstWallet, amount)
}
