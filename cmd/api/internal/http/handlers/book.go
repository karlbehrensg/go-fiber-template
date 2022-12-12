package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/internal/service"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"github.com/karlbehrensg/go-fiber-template/pkg/utils"
)

type BookHandler struct {
	Service service.BookService
}

// CreateBook godoc
// @Summary create book.
// @Description endpoint for create books.
// @Tags Book
// @Accept json
// @Produce json
// @Param Body body requests.BookRequest true "The body to book"
// @Success 201 {object} responses.BookResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /book [post]
// CreateBook controller to create book
func (h BookHandler) CreateBook(c *fiber.Ctx) error {
	// Convert the request data to the structure
	data := &requests.BookRequest{}
	if err := c.BodyParser(&data); err != nil {
		logger.Error("Error decode json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	// validates the structure
	if err := utils.GetValidator().Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// calls use case to create book
	if err := h.Service.CreateBook(*data); err != nil {
		return c.Status(err.Code).JSON(err.AsMessage())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created",
	})
}

// GetAllBook godoc
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {array} responses.BookResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /book [get]
// GetAllAuthor controller to get all exists book
func (h BookHandler) GetAllBook(c *fiber.Ctx) error {
	var response []responses.BookResponse
	var err *errs.AppError

	// calls use case to get all book
	if response, err = h.Service.FindAllBook(); err != nil {
		return c.Status(err.Code).JSON(err.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetBookById godoc
// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path integer true "Book ID"
// @Success 200 {object} responses.BookResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /book/{id} [get]
// GetBookById controller to find book by ID
func (h BookHandler) GetBookById(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Book id",
		})
	}

	var response *responses.BookResponse
	var appErr *errs.AppError
	// calls use case to find book by ID
	if response, appErr = h.Service.FindBookById(uint(id)); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateBook godoc
// @Summary update book.
// @Description endpoint for update books.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path integer true "Book ID"
// @Param Body body requests.BookRequest true "The body to book"
// @Success 200 {object} responses.BookResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /book/{id} [put]
// UpdateBook controller to update book
func (h BookHandler) UpdateBook(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Book id",
		})
	}

	// Convert the request data to the structure
	data := &requests.BookRequest{}
	data.Id = uint(id)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	// validates the structure
	if err := utils.GetValidator().Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var response *responses.BookResponse
	var appErr *errs.AppError
	// calls use case to update book
	if response, appErr = h.Service.UpdateBook(data); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteBook godoc
// @Summary delete book.
// @Description endpoint for delete books.
// @Tags Book
// @Accept json
// @Produce json
// @Param id path integer true "Book ID"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /book/{id} [delete]
// DeleteBook controller to delete book
func (h BookHandler) DeleteBook(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Book id",
		})
	}

	var appErr *errs.AppError
	// calls use case to delete author
	if appErr = h.Service.DeleteBook(uint(id)); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Book deleted",
	})
}
