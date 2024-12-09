package main

import (
	"savioafs/daily-diet-app-go/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	_, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	server.Run(":8080")
}
