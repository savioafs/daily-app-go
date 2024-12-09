package main

import (
	"savioafs/daily-diet-app-go/internal/config"
)

func main() {

	dbConn, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	server := config.SetupRoutes(dbConn)

	server.Run(":8080")
}
