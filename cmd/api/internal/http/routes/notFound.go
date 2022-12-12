package routes

import "github.com/gofiber/fiber/v2"

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(router *fiber.App) {
	// Register new special route.
	router.Use(
		// Anonimus function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "sorry, endpoint is not found",
			})
		},
	)
}
