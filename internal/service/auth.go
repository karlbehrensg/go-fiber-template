package service

import (
	"strings"

	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
)

type AuthService interface {
	Login(requests.LoginRequest) (*responses.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.UserRepository
}

// NewAuthService create a new instance of DefaultAuthService
func NewAuthService(repository domain.UserRepository) DefaultAuthService {
	return DefaultAuthService{repository}
}

// Login use case for validate user and return token
func (s DefaultAuthService) Login(request requests.LoginRequest) (*responses.LoginResponse, *errs.AppError) {
	var appErr *errs.AppError
	var u *domain.User

	// find user by email
	if u, appErr = s.repo.FindUserByEmail(request.Email); appErr != nil {
		logger.Error(appErr.Message)
		if strings.Contains(appErr.Message, "record not found") {
			return nil, errs.NewAuthenticationError("invalid credentials")
		}
		return nil, appErr
	}

	// validate password
	if appErr = u.ComparePassword(request.Password); appErr != nil {
		return nil, appErr
	}

	// create claim
	jwtClaims := u.ToNewUtilsJWTClaims()

	var accessToken string
	var err error
	// create token
	if accessToken, err = jwtClaims.CreateToken(); err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected error while creating token")
	}

	return &responses.LoginResponse{Token: accessToken}, nil
}
