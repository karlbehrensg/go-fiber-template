package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"github.com/karlbehrensg/go-fiber-template/pkg/utils"
)

// ValidateJWT middleware to validate JWT
func ValidateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed token",
			})
		}

		token := strings.TrimSpace(strings.Split(authHeader, " ")[1])
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed token",
			})
		}

		// Validate token
		claims := &utils.JWTClaims{}
		if err := claims.ValidateToken(token); err != nil {
			logger.Error(err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
			})
		}

		return c.Next()
	}
}
