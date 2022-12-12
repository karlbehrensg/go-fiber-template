package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/internal/service"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"github.com/karlbehrensg/go-fiber-template/pkg/utils"
)

type AuthHandler struct {
	UserSrv service.UserService
	AuthSrv service.AuthService
}

// SignUp godoc
// @Summary register users.
// @Description endpoint for register user.
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param full_name formData string false "Full name"
// @Param email formData string true "email"
// @Param email_confirmation formData string true "email"
// @Param password formData string true "password"
// @Param password_confirmation formData string true "password"
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /auth/signup [post]
// SignUp controller to register users
func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	// Convert the request data to the structure
	data := &requests.UserRequest{}
	if err := c.BodyParser(data); err != nil {
		logger.Error(fmt.Sprintf("Error decode: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid data",
		})
	}

	// validates the structure
	if err := utils.GetValidator().Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// calls use case to create user
	if err := h.UserSrv.CreateUser(*data); err != nil {
		return c.Status(err.Code).JSON(err.AsMessage())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}

// Login godoc
// @Summary create token.
// @Description endpoint for get token.
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 201 {object} responses.LoginResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /auth/login [post]
// Login controller to generate token
func (h AuthHandler) Login(c *fiber.Ctx) error {
	// Convert the request data to the structure
	data := requests.LoginRequest{}
	if err := c.BodyParser(&data); err != nil {
		logger.Error(fmt.Sprintf("Error decode: %s", err.Error()))
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

	var response *responses.LoginResponse
	var appErr *errs.AppError
	// calls use case to generate token
	if response, appErr = h.AuthSrv.Login(data); appErr != nil {
		return c.Status(appErr.Code).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
