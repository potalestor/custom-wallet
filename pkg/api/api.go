package api

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/potalestor/custom-wallet/api/docs"
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

// Build creates
// @title CUSTOM-WALLET REST API
// @version 0.0.1
// @description Swagger API for Golang Project CUSTOM-WALLET.
// @contact.name @potalestor
// @contact.email potalestor@gmail.com
// @BasePath /api/v1
func (a *API) Build() *gin.Engine {
	r := gin.New()
	r.Use(Recovery, Logger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r)

	v1 := r.Group("/v1/api")

	a.buildWalletApi(v1.Group("/wallets")) //.
	// 	buildTransferApi(v1.Group("/transfer")).
	// 	buildReportApi(v1.Group("/reoirts"))

	return r
}

func (a *API) buildWalletApi(r *gin.RouterGroup) *API {
	r.POST("/", a.Wallet.CreateWallet)
	return a
}

// func (a *API) buildOrgApi(r *gin.RouterGroup) *API {
// 	r.GET("/", a.Org.Orgs)
// 	r.GET("/:org", a.Org.Org)
// 	r.POST("/:org/", a.Org.Add)
// 	r.GET("/:org/suspension", a.Org.Suspended)
// 	r.POST("/:org/suspension", a.Org.SuspendOn)
// 	r.DELETE("/:org/suspension", a.Org.SuspendOff)
// 	r.GET("/:org/profile", a.Org.GetProfile)
// 	r.PUT("/:org/profile", a.Org.AddProfile)

// 	return a
// }

// func (a *API) buildUserApi(r *gin.RouterGroup) *API {
// 	r.GET("/", a.User.Users)
// 	r.POST("/", a.User.Add)
// 	r.GET("/:userid/suspension", a.User.Suspended)
// 	r.POST("/:userid/suspension", a.User.SuspendOn)
// 	r.DELETE("/:userid/suspension", a.User.SuspendOff)
// 	r.POST("/:userid/active", a.User.SetActive)
// 	r.GET("/:userid/profile", a.User.GetProfile)
// 	r.PUT("/:userid/profile", a.User.AddProfile)
// 	r.GET("/:userid/keys/", a.User.Keys)
// 	r.PUT("/:userid/certificate", a.User.AddCert)
// 	r.DELETE("/:userid/certificates/:certid", a.User.DelCert)

// 	return a
// }

// func (a *API) buildDocumentApi(r *gin.RouterGroup) *API {
// 	r.GET("/:template/", a.Document.Docs)
// 	r.POST("/:template/", a.Document.Add)
// 	r.PUT("/:template/:docid", a.Document.Update)
// 	r.GET("/:template/:docid", a.Document.Doc)
// 	r.GET("/:template/:docid/parties", a.Document.Parties)
// 	r.DELETE("/:template/:docid", a.Document.Del)

// 	return a
// }
