package main

import (
	"log"
	"os"

	db "fiber-backend/database"
	r "fiber-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New() // Initialize fiber instance

	if err := godotenv.Load(); err != nil { // environment variable
		panic(err)
	}

	app.Use(logger.New()) // Logger middleware

	db.MongoConnect() // MongoDB connection

	r.SetupRoutes(app) // Initialize routes

	log.Fatal(app.Listen(os.Getenv("APP_PORT"))) // listening server on port: 3000

}




















