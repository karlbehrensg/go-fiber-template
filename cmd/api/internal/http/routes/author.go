package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/handlers"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/middlewares"
	"github.com/karlbehrensg/go-fiber-template/internal/repository"
	"github.com/karlbehrensg/go-fiber-template/internal/service"

	"gorm.io/gorm"
)

// AuthorRoutes endpoints for the author section
func AuthorRoutes(router *fiber.App, dbClient *gorm.DB) {
	h := handlers.AuthorHandler{
		Service: service.NewAuthorService(repository.NewAuthorRepositoryGorm(dbClient)),
	}
	api := router.Group("/author")
	api.Use(middlewares.ValidateJWT())
	api.Post("", h.CreateAuthor)
	api.Get("", h.GetAllAuthor)
	api.Get("/:id", h.GetAuthorById)
	api.Put("/:id", h.UpdateAuthor)
	api.Delete("/:id", h.DeleteAuthor)
}
