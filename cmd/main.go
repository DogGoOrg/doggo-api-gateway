package main

import (
	"log"
	"os"

	"github.com/DogGoOrg/doggo-api-gateway/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	//config
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//logging
	logger := server.SetupLogger()

	app := server.NewServer(logger)

	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}

}
