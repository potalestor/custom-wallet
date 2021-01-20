package api

import (
	"log"
	"net/http/httputil"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	formatRecovery = `[Recovery] panic:\n%s\n%s\n%s`
	bit16          = 1 << 16
)

// Recovery middleware catches a panic and handles it.
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httpRequest, _ := httputil.DumpRequest(c.Request, false)

			buf := make([]byte, bit16)
			stackSize := runtime.Stack(buf, true)

			log.Printf(formatRecovery, string(httpRequest), err, string(buf[0:stackSize]))
		}
	}()

	c.Next()
}
