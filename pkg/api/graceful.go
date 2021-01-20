package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "unsafe" // for resolveAddress

	_ "github.com/gin-gonic/gin" // for resolveAddress
)

var stopServerExpiration = 20 * time.Second

// Graceful server is used for safely shutting down server.
type Graceful struct {
	handler http.Handler
}

// NewGraceful creates graceful server.
func NewGraceful(handler http.Handler) *Graceful {
	return &Graceful{
		handler: handler,
	}
}

// Run server and wait for it to finish.
func (s *Graceful) Run(addr ...string) error {
	srv := &http.Server{
		Addr:    resolveAddress(addr),
		Handler: s.handler,
	}

	log.Printf("server is starting on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	// kill -2 is syscall.SIGINT
	// kill (no param) default send syscall.SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	log.Printf("server is shutdowning: '%v' signal has been received", sig)

	// The context is used to inform the server it has 20 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), stopServerExpiration)
	defer cancel()

	return srv.Shutdown(ctx)
}

//go:linkname resolveAddress github.com/gin-gonic/gin.resolveAddress
func resolveAddress(addr []string) string
