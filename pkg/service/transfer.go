package service

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
	"github.com/potalestor/custom-wallet/pkg/repo"
)

type Transfer struct {
	repository repo.Repository
}

func NewTransfer(repository repo.Repository) *Transfer {
	return &Transfer{repository: repository}
}

func (t *Transfer) Transfer(src, dst string, amount model.USD) error {
	if err := amount.Validate(); err != nil {
		return err
	}
	
	ctx := context.Background()

	srcWallet := model.Wallet{Name: src}

	if err := t.repository.GetWalletByName(ctx, &srcWallet); err != nil {
		return err
	}

	dstWallet := model.Wallet{Name: dst}

	if err := t.repository.GetWalletByName(ctx, &dstWallet); err != nil {
		return err
	}

	return t.repository.Transfer(ctx, &srcWallet, &dstWallet, amount)
}
