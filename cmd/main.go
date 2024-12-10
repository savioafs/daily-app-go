package main

import (
	"savioafs/daily-diet-app-go/internal/config"
)

func main() {

	dbConn, jwtAuth, err := config.LoadConfigs()
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	server := config.SetupRoutes(dbConn, jwtAuth)

	server.Run(":8080")
}
