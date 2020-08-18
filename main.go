package main

import (
	"log"

	"github.com/firebase007/go-rest-api-with-fiber/database"
	"github.com/firebase007/go-rest-api-with-fiber/router"
	"github.com/gofiber/fiber" // import the fiber package
	"github.com/gofiber/fiber/middleware"

	_ "github.com/lib/pq"
)

// entry point to our program
func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Middleware
	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(3000)

}
