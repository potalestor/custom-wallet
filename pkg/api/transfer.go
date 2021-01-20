package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/potalestor/custom-wallet/pkg/app"
	"github.com/potalestor/custom-wallet/pkg/model"
)

type Transfer struct {
	wallet *app.Wallet
}

func NewTransfer(wallet *app.Wallet) *Transfer {
	return &Transfer{wallet: wallet}
}

// Transfer money from one wallet to another.
// @Summary Transfer money from one wallet to another.
// @Description Transfer money from one wallet to another.
// @Tags transfers
// @Produce json
// @Param  src_wallet path string true "source wallet name"
// @Param  dst_wallet path string true "destination wallet name"
// @Param  amount path number true "amount"
// @Success 200 {array} model.Wallet
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /transfers/{src_wallet}/{dst_wallet}/{amount} [put]
func (t *Transfer) Transfer(c *gin.Context) {
	srcWallet := c.Param("src_wallet")
	if err := validation.Validate(srcWallet,
		validation.Required,
		is.Alphanumeric,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	dstWallet := c.Param("dst_wallet")
	if err := validation.Validate(dstWallet,
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

	result, err := t.wallet.Transfer.Transfer(srcWallet, dstWallet, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, result)
}
