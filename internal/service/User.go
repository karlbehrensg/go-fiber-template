package service

import (
	"fmt"

	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
)

// UserService port primary
type UserService interface {
	CreateUser(requests.UserRequest) *errs.AppError
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

func (s DefaultUserService) CreateUser(request requests.UserRequest) *errs.AppError {
	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := user.HashPassword(); err != nil {
		logger.Error(fmt.Sprintf("Error while encripting password: %s", err.Error()))
		return errs.NewUnexpectedError("Unexpected error while encripting password")
	}

	if err := s.repo.SaveUser(user); err != nil {
		return err
	}

	return nil
}
