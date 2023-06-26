package main

import (
	"log"
	"os"

	"github.com/DogGoOrg/doggo-api-gateway/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.NewServer()
	app.Engine().Use(gin.Recovery())

	conns := server.EstablishConnection()
	server.ConfigureRoutes(app, conns)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
