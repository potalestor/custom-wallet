package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter is a response writer to capture the response body.
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)

	return r.ResponseWriter.Write(b)
}

// logParams is the structure any formatter will be handed when time to log comes.
type logParams struct {
	gin.LogFormatterParams
}

func (params *logParams) String() string {
	return fmt.Sprintf("[GIN] | %3d | %13v | %15s | %-7s %#v %s",
		params.StatusCode,
		params.Latency,
		params.ClientIP,
		params.Method,
		params.Path,
		params.ErrorMessage,
	)
}

// logHandler logs request to the application logger.
func (params *logParams) logHandler() {
	code := params.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		log.Println("[INFO ]", params.String())
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		log.Println("[DEBUG]", params.String())
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		log.Println("[WARN ]", params.String())
	default:
		log.Println("[ERROR]", params.String())
	}
}

// Logger middleware logs http and other requests in the application logger.
func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	c.Next()

	if raw != "" {
		path = path + "?" + raw
	}

	err := ""
	if c.Writer.Status() >= http.StatusBadRequest {
		err = c.Errors.ByType(gin.ErrorTypePrivate).String()
		if (w.body.Len() > 0) &&
			(c.Errors.ByType(gin.ErrorTypePrivate) == nil) {
			err = w.body.String()
		}
	}

	params := logParams{
		gin.LogFormatterParams{
			Request:      c.Request,
			StatusCode:   c.Writer.Status(),
			Latency:      time.Since(start),
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			Path:         path,
			ErrorMessage: err,
			BodySize:     c.Writer.Size(),
			Keys:         c.Keys,
		},
	}
	params.logHandler()
}
