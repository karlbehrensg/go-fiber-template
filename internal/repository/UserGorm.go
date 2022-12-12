package repository

import (
	"strings"

	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	client *gorm.DB
}

// NewUserRepositoryGorm create a new instance of UserRepositoryGorm
func NewUserRepositoryGorm(dbClient *gorm.DB) UserRepositoryGorm {
	return UserRepositoryGorm{dbClient}
}

// SaveUser save user in database
func (r UserRepositoryGorm) SaveUser(user *domain.User) *errs.AppError {
	if err := r.client.Create(user).Error; err != nil {
		logger.Error(err.Error())
		if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint ") {
			return errs.NewBadRequestError("key email duplicate value")
		}
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}

// FindUserByEmail find user by Email in database
func (r UserRepositoryGorm) FindUserByEmail(email string) (*domain.User, *errs.AppError) {
	var user *domain.User

	if err := r.client.Where("email = ?", email).First(&user).Error; err != nil {
		logger.Error(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return user, nil
}
