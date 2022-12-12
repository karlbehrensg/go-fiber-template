package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SwaggerRoutes func for describe group of API Docs routes.
func SwaggerRoutes(router *fiber.App) {
	// Create routes group.
	api := router.Group("/docs")
	api.Get("*", swagger.HandlerDefault)
}
