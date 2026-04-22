package main

import (
	"fmt"
	"server/internal/routes"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   🖥️  REMOTE CONTROL SERVER - PC       ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Printf("\n📡 IP: %s\n", utils.GetLocalIP())
	fmt.Println("🔌 Port: 8080")
	fmt.Println("\n📱 Connect the app to this IP")
	fmt.Println("⚡ Press Ctrl+C to stop")

	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.0.0/16"})

	routes.Setup(r)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
}
