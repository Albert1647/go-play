package main

import (
	"github.com/gin-gonic/gin"
	"natthan.com/go-play/db"
	"natthan.com/go-play/routes"
)

func main() {
	// Prepare Table if not exist
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
