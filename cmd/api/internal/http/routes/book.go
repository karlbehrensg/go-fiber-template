package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/handlers"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/middlewares"
	"github.com/karlbehrensg/go-fiber-template/internal/repository"
	"github.com/karlbehrensg/go-fiber-template/internal/service"

	"gorm.io/gorm"
)

// BookRoutes endpoints for the book section
func BookRoutes(router *fiber.App, dbClient *gorm.DB) {
	h := handlers.BookHandler{
		Service: service.NewBookService(repository.NewBookRepositoryGorm(dbClient)),
	}
	// Create routes group.
	api := router.Group("/book")
	// use middleware for validate JWT
	api.Use(middlewares.ValidateJWT())
	api.Post("", h.CreateBook)
	api.Get("", h.GetAllBook)
	api.Get("/:id", h.GetBookById)
	api.Put("/:id", h.UpdateBook)
	api.Delete("/:id", h.DeleteBook)
}
