package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/potalestor/custom-wallet/pkg/app"
	"github.com/potalestor/custom-wallet/pkg/model"
)

// Report handler.
type Report struct {
	wallet *app.Wallet
}

// NewReport returns new instance.
func NewReport(wallet *app.Wallet) *Report {
	return &Report{wallet: wallet}
}

// Report on the wallet.
// @Summary Report on the wallet.
// @Description Report on the wallet. Using Filter. Operation: 1-Deposit, 2-Withdraw, 3-Both. Date range using RFC3339.
// @Tags reports
// @Param filter body model.Filter true "Create Filter"
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /reports [put].
func (r *Report) Report(c *gin.Context) {
	var filter model.Filter

	if err := c.ShouldBindBodyWith(&filter, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := filter.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	reports, err := r.wallet.Report.Report(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	if err := reports.CSV(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Status(http.StatusOK)
}
