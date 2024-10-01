package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndkode/elabram-backend-recruitment/cmd/configs"
)

func main() {
	r := gin.Default()

	// Connect to Database
	configs.ConnectDB()

	// Run Server
	r.Run(":8080")
}
