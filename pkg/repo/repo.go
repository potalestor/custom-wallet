package repo

import (
	"context"

	"github.com/potalestor/custom-wallet/pkg/model"
)

// Repository contains database operations.
type Repository interface {
	Open() error
	Close() error
	CreateWallet(ctx context.Context, wallet *model.Wallet) error
	GetWalletByName(ctx context.Context, wallet *model.Wallet) error
	Deposit(ctx context.Context, wallet *model.Wallet, amount model.USD) error
	Transfer(ctx context.Context, src, dst *model.Wallet, amount model.USD) error
}
