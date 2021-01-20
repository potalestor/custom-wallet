package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/potalestor/custom-wallet/pkg/api/docs" // for swagger documentation
	"github.com/potalestor/custom-wallet/pkg/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// API Web-service.
type API struct {
	Wallet   *Wallet
	Transfer *Transfer
	Report   *Report
}

// NewAPI returns new instance.
func NewAPI(wallet *app.Wallet) *API {
	return &API{
		Wallet:   NewWallet(wallet),
		Transfer: NewTransfer(wallet),
		Report:   NewReport(wallet),
	}
}

// Build creates.
// @title CUSTOM-WALLET REST API
// @version 0.0.1
// @description Swagger API for Golang Project CUSTOM-WALLET.
// @contact.name potalestor@gmail.com
// @BasePath /api/v1.
func (a *API) Build() *gin.Engine {
	r := gin.New()
	r.Use(Recovery, Logger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r)

	v1 := r.Group("/api/v1")

	a.buildWalletAPI(v1.Group("/wallets")).
		buildTransferAPI(v1.Group("/transfers")).
		buildReportAPI(v1.Group("/reports"))

	return r
}

func (a *API) buildWalletAPI(r *gin.RouterGroup) *API {
	r.POST("/:wallet_name", a.Wallet.Create)
	r.PUT("/:wallet_name/:amount", a.Wallet.Deposit)

	return a
}

func (a *API) buildTransferAPI(r *gin.RouterGroup) *API {
	r.PUT("/:src_wallet/:dst_wallet/:amount", a.Transfer.Transfer)

	return a
}

func (a *API) buildReportAPI(r *gin.RouterGroup) *API {
	r.PUT("/", a.Report.Report)

	return a
}
