package api

import (
	"fmt"
	"log"
	"net/http/httputil"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// Recovery middleware catches a panic and handles it.
func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			headers := strings.Split(string(httpRequest), "\r\n")

			buf := make([]byte, 1<<16)
			stackSize := runtime.Stack(buf, true)

			log.Printf(fmt.Sprintf("[Recovery] panic:\n%s\n%s\n%s",
				strings.Join(headers, "\r\n"),
				err,
				string(buf[0:stackSize])))
		}
	}()

	c.Next()
}
