package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/potalestor/custom-wallet/pkg/app"
	"github.com/potalestor/custom-wallet/pkg/model"
)

type Wallet struct {
	wallet *app.Wallet
}

func NewWallet(wallet *app.Wallet) *Wallet {
	return &Wallet{wallet: wallet}
}

// Create new wallet.
// @Summary Create new wallet.
// @Description Create new wallet.
// @Tags wallets
// @Produce json
// @Param  wallet_name path string true "wallet name"
// @Success 200 {object} model.Wallet
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /wallets/{wallet_name} [post]
func (w *Wallet) Create(c *gin.Context) {
	walletName := c.Param("wallet_name")
	if err := validation.Validate(walletName,
		validation.Required,
		is.Alphanumeric,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	result, err := w.wallet.Wallet.CreateWallet(walletName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, result)
}

// Deposit money to wallet.
// @Summary Deposit money to wallet.
// @Description Deposit money to wallet.
// @Tags wallets
// @Produce json
// @Param  wallet_name path string true "wallet name"
// @Param  amount path number true "amount"
// @Success 200 {object} model.Wallet
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /wallets/{wallet_name}/{amount} [put]
func (w *Wallet) Deposit(c *gin.Context) {
	walletName := c.Param("wallet_name")
	if err := validation.Validate(walletName,
		validation.Required,
		is.Alphanumeric,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	var amount model.USD
	if err := amount.Parse(c.Param("amount")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	result, err := w.wallet.Wallet.Deposit(walletName, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, result)
}
