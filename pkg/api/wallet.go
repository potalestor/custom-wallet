package api

import (
	"github.com/gin-gonic/gin"
	"github.com/potalestor/custom-wallet/pkg/app"
)

type Wallet struct {
	wallet *app.Wallet
}

func NewWallet(wallet *app.Wallet) *Wallet {
	return &Wallet{wallet: wallet}
}

// CreateWallet returns new wallet.
// @Summary Create new wallet.
// @Description Create new wallet.
// @Tags wallets
// @Param  name path string true "wallet name"
// @Success 200
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /wallets/{name} [post]
func (w *Wallet) CreateWallet(c *gin.Context) {

}

// func (a *Auth) Authenticate(c *gin.Context) {
// 	addr, err := a.tf.Auth.Autenticate(c.GetHeader("Authorization"))
// 	if err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusUnauthorized, err)
// 		return
// 	}

// 	claims := claim.NewBuilder().
// 		BuildAddress(addr).
// 		BuildClientIP(c.ClientIP()).
// 		BuildSessionID(SessionID(c)).
// 		Build()

// 	c.Set("claims", claims)

// 	c.Next()
// }

// func (a *Auth) Login(c *gin.Context) {
// 	claims, ok := c.Value("claims").(*claim.Claims)
// 	if !ok {
// 		a.respondErrorWithChallenge(c, http.StatusInternalServerError, tferr.ErrClaimsRequired)
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := a.tf.Auth.Login(ctx, claims); err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusUnauthorized, err)
// 		return
// 	}

// 	token, err := claims.Token()
// 	if err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, token)
// }

// func (a *Auth) Logout(c *gin.Context) {
// 	claims, ok := c.Value("claims").(*claim.Claims)
// 	if !ok {
// 		a.respondErrorWithChallenge(c, http.StatusInternalServerError, tferr.ErrClaimsRequired)
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := a.tf.Auth.Logout(ctx, claims); err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// func (a *Auth) Authorize(c *gin.Context) {
// 	claims := claim.NewClaims()

// 	if err := claims.Parse(c.Request.Header.Get("Authorization")); err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusUnauthorized, err)
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := a.tf.Auth.Authorize(ctx, claims); err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusUnauthorized, err)
// 		return
// 	}

// 	c.Set("claims", claims)

// 	c.Next()
// }

// func (a *Auth) Refresh(c *gin.Context) {
// 	claims, ok := c.Value("claims").(*claim.Claims)
// 	if !ok {
// 		a.respondErrorWithChallenge(c, http.StatusInternalServerError, tferr.ErrClaimsRequired)
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	claims, err := a.tf.Auth.Refresh(ctx, claims)
// 	if err != nil {
// 		a.respondErrorWithChallenge(c, http.StatusUnauthorized, err)
// 		return
// 	}

// 	c.Set("claims", claims)
// }

// func (a *Auth) respondErrorWithChallenge(c *gin.Context, code int, message interface{}) {
// 	token, err := a.tf.Auth.CreateChallenge()
// 	if err != nil {
// 		SpanLogger(c).Error(err.Error())
// 	} else {
// 		newauth := fmt.Sprintf("Newauth realm=\"cbg-request\", challenge=%q", token)
// 		c.Header("WWW-Authenticate", newauth)
// 	}

// 	c.AbortWithStatusJSON(code, gin.H{"error": message})
// }
