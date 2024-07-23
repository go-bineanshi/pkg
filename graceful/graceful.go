package graceful

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulServer holds the Gin Engine and the HTTP server
type GracefulServer struct {
	Engine *gin.Engine
	Server *http.Server
}

// NewGracefulServer creates a new instance of GracefulServer
func NewGracefulServer(addr string, engine *gin.Engine) *GracefulServer {
	server := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	return &GracefulServer{
		Engine: engine,
		Server: server,
	}
}

// Start runs the server with graceful shutdown capabilities
func (gs *GracefulServer) Start() {
	go func() {
		if err := gs.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := gs.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
