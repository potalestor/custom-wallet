package api

import (
	_ "github.com/potalestor/custom-wallet/pkg/api/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/potalestor/custom-wallet/pkg/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service interface {
	Build() *gin.Engine
}

type API struct {
	Wallet   *Wallet
	Transfer *Transfer
	Report   *Report
}

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
// @BasePath /api/v1
func (a *API) Build() *gin.Engine {
	r := gin.New()
	r.Use(Recovery, Logger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r)

	v1 := r.Group("/api/v1")

	a.buildWalletApi(v1.Group("/wallets")).
		buildTransferApi(v1.Group("/transfers")).
		buildReportApi(v1.Group("/reports"))

	return r
}

func (a *API) buildWalletApi(r *gin.RouterGroup) *API {
	r.POST("/:wallet_name", a.Wallet.Create)
	r.PUT("/:wallet_name/:amount", a.Wallet.Deposit)

	return a
}

func (a *API) buildTransferApi(r *gin.RouterGroup) *API {
	r.PUT("/:src_wallet/:dst_wallet/:amount", a.Transfer.Transfer)

	return a
}

func (a *API) buildReportApi(r *gin.RouterGroup) *API {
	r.PUT("/", a.Report.Report)

	return a
}
