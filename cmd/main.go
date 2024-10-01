package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ndkode/elabram-backend-recruitment/cmd/configs"
	"github.com/ndkode/elabram-backend-recruitment/cmd/routes"
)

func main() {
	r := gin.Default()

	// Gzip Compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Connect to Database
	configs.ConnectDB()

	// Setup Router
	routes.SetupRouter(r)

	// Run Server
	r.Run(":8080")
}
