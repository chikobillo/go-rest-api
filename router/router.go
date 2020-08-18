package router

import (
	"github.com/chikobillo/go-rest-api/handler"
	"github.com/chikobillo/go-rest-api/middleware"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), middleware.AuthReq())

	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
