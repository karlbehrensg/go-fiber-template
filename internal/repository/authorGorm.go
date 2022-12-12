package repository

import (
	"fmt"
	"strings"

	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"gorm.io/gorm"
)

type AuthorRepositoryGorm struct {
	client *gorm.DB
}

// NewAuthorRepositoryGorm create a new instance of AuthorRepositoryGorm
func NewAuthorRepositoryGorm(dbClient *gorm.DB) AuthorRepositoryGorm {
	return AuthorRepositoryGorm{dbClient}
}

// SaveAuthor save author in database
func (r AuthorRepositoryGorm) SaveAuthor(author *domain.Author) *errs.AppError {
	if err := r.client.Create(author).Error; err != nil {
		logger.Error(err.Error())
		if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint ") {
			return errs.NewUnexpectedError("key full_name duplicate value")
		}
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}

// FindAllAuthor find all author in database
func (r AuthorRepositoryGorm) FindAllAuthor() ([]domain.Author, *errs.AppError) {
	authors := []domain.Author{}
	if err := r.client.Find(&authors).Error; err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return authors, nil
}

// FindAuthorById find author by ID in database
func (r AuthorRepositoryGorm) FindAuthorById(id uint) (*domain.Author, *errs.AppError) {
	var author *domain.Author

	if err := r.client.Where("id = ?", id).First(&author).Error; err != nil {
		logger.Error(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return author, nil
}

// UpdateAuthor update author in database
func (r AuthorRepositoryGorm) UpdateAuthor(author *domain.Author) (*domain.Author, *errs.AppError) {
	var result *gorm.DB
	if result = r.client.Where("id = ?", author.ID).Updates(&author); result.Error != nil {
		logger.Error(result.Error.Error())
		if strings.Contains(result.Error.Error(), "ERROR: duplicate key value violates unique constraint ") {
			return nil, errs.NewUnexpectedError("key full_name duplicate value")
		}
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// validates if the rows have changed
	if result.RowsAffected < 1 {
		logger.Info(fmt.Sprintf("Row with id=%d cannot be updated because it doesn't exist", author.ID))
		return nil, errs.NewNotFoundError("Author not found")
	}

	return author, nil
}

// DeleteAuthor delete author in database
func (r AuthorRepositoryGorm) DeleteAuthor(id uint) *errs.AppError {
	var author *domain.Author
	var result *gorm.DB
	if result = r.client.Where("id = ?", id).Delete(&author); result.Error != nil {
		logger.Error(result.Error.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	// validates if the rows have changed
	if result.RowsAffected < 1 {
		logger.Info(fmt.Sprintf("Row with id=%d cannot be deleted because it doesn't exist", id))
		return errs.NewNotFoundError("Author not found")
	}

	return nil
}
