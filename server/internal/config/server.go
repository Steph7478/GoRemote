package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"server/internal/routes"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ginEngine         *gin.Engine
	httpServer        *http.Server
	activeConnections int
)

func startServer() {
	gin.SetMode(gin.ReleaseMode)

	ginEngine = gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16", "10.0.0.0/8"})

	ginEngine.Use(func(c *gin.Context) {
		local := c.ClientIP() == "127.0.0.1" || c.ClientIP() == "::1"

		if !local {
			activeConnections++
			fmt.Printf("📊 Active connections: %d\n", activeConnections)
		}

		c.Next()

		if !local {
			activeConnections--
			fmt.Printf("📊 Active connections: %d\n", activeConnections)

			if activeConnections == 0 {
				time.AfterFunc(10*time.Second, func() {
					if activeConnections == 0 && serverRunning {
						closeServerPort()
					}
				})
			}
		}
	})

	routes.Setup(ginEngine)
	fmt.Println("🔒 Server created, waiting for discovery")
}

func openServerPort() {
	if serverRunning {
		return
	}

	httpServer = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: ginEngine,
	}

	serverRunning = true

	go func() {
		fmt.Println("🔓 Discovery found! Opening port 8080...")

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("❌ Error: %v\n", err)
			serverRunning = false
		}
	}()
}

func closeServerPort() {
	if !serverRunning || httpServer == nil {
		return
	}

	serverRunning = false
	fmt.Println("🔒 Closing port 8080...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("❌ Error closing port: %v\n", err)
		return
	}

	httpServer = nil
	fmt.Println("🔒 Port 8080 closed")
}
