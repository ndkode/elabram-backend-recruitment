package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Run Server
	r.Run(":8080")
}
