package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/handlers"
	"github.com/karlbehrensg/go-fiber-template/internal/repository"
	"github.com/karlbehrensg/go-fiber-template/internal/service"

	"gorm.io/gorm"
)

// AuthRoutes endpoints to authentication
func AuthRoutes(router *fiber.App, dbClient *gorm.DB) {
	c := handlers.AuthHandler{
		UserSrv: service.NewUserService(repository.NewUserRepositoryGorm(dbClient)),
		AuthSrv: service.NewAuthService(repository.NewUserRepositoryGorm(dbClient)),
	}
	api := router.Group("/auth")
	api.Post("/signup", c.SignUp)
	api.Post("login", c.Login)
}
