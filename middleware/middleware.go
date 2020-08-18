package middleware

import (
	"github.com/chikobillo/go-rest-api/config"
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
)

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}
