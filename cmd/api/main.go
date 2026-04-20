package main

import (
	"net/http"
	"os"

	"github.com/doteneff/bijakbudget-api/internal/config"
	"github.com/doteneff/bijakbudget-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Connect database
	config.ConnectDB()

	// 2. Setup Gin Router
	r := gin.Default()

	// 3. Define basic routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from bijakbudget-api",
		})
	})

	// 4. Setup API Routing
	routes.SetupRouter(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}