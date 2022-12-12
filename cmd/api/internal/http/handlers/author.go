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

type AuthorHandler struct {
	Service service.AuthorService
}

// CreateAuthor godoc
// @Summary create author.
// @Description endpoint for create authors.
// @Tags Author
// @Accept json
// @Produce json
// @Param Body body requests.AuthorRequest true "The body to author"
// @Success 201 {object} responses.AuthorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /author [post]
// CreateAuthor controller to create author
func (h AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	// Convert the request data to the structure
	data := &requests.AuthorRequest{}
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

	// calls use case to create author
	if err := h.Service.CreateAuthor(*data); err != nil {
		return c.Status(err.Code).JSON(err.AsMessage())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Author created",
	})
}

// GetAllAuthor godoc
// @Description Get all exists authors.
// @Summary get all exists authors
// @Tags Author
// @Accept json
// @Produce json
// @Success 200 {array} responses.AuthorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /author [get]
// GetAllAuthor controller to get all author
func (h AuthorHandler) GetAllAuthor(c *fiber.Ctx) error {
	var response []responses.AuthorResponse
	var err *errs.AppError
	// calls use case to get all author
	if response, err = h.Service.FindAllAuthor(); err != nil {
		return c.Status(err.Code).JSON(err.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAuthorById godoc
// @Description Get author by given ID.
// @Summary get author by given ID
// @Tags Author
// @Accept json
// @Produce json
// @Param id path integer true "Author ID"
// @Success 200 {object} responses.AuthorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /author/{id} [get]
// GetAuthorById controller to find author by ID
func (h AuthorHandler) GetAuthorById(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author id",
		})
	}

	var response *responses.AuthorResponse
	var appErr *errs.AppError
	// calls use case to find author by ID
	if response, appErr = h.Service.FindAuthorById(uint(id)); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateAuthor godoc
// @Summary update author.
// @Description endpoint for update authors.
// @Tags Author
// @Accept json
// @Produce json
// @Param id path integer true "Author ID"
// @Param Body body requests.AuthorRequest true "The body to author"
// @Success 200 {object} responses.AuthorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /author/{id} [put]
// UpdateAuthor controller to update author
func (h AuthorHandler) UpdateAuthor(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author id",
		})
	}

	// Convert the request data to the structure
	data := &requests.AuthorRequest{}
	data.Id = uint(id)
	if err := c.BodyParser(&data); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	// validates the structure
	if err := utils.GetValidator().Struct(data); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var response *responses.AuthorResponse
	var appErr *errs.AppError
	// calls use case to update author
	if response, appErr = h.Service.UpdateAuthor(data); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteAuthor godoc
// @Summary delete author.
// @Description endpoint for delete authors.
// @Tags Author
// @Accept json
// @Produce json
// @Param id path integer true "Author ID"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security Bearer
// @Router /author/{id} [delete]
// DeleteAuthor controller to delete author
func (h AuthorHandler) DeleteAuthor(c *fiber.Ctx) error {
	var id int
	var err error
	// get ID parameter from url
	if id, err = c.ParamsInt("id"); err != nil {
		logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author id",
		})
	}

	var appErr *errs.AppError
	// calls use case to delete author
	if appErr = h.Service.DeleteAuthor(uint(id)); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Author deleted",
	})
}
