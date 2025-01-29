package main

import (
	"library-api-author/internal/config"
	"library-api-author/internal/factory"
	"library-api-author/internal/routes"
	"library-api-author/pkg/database"
	"log"
)

func main() {
	config.LoadConfig()
	psqlDB, err := database.NewPqSQLClient()
	if err != nil {
		log.Fatal("Could not connect to MySQL:", err)
	}

	provider := factory.InitFactory(psqlDB)

	router := routes.RegisterRoutes(provider)
	log.Printf("Server running on :%s\n", config.ENV.ServerPort)
	log.Fatal(router.Run(":" + config.ENV.ServerPort))
}
