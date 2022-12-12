package repository

import (
	"fmt"
	"strings"

	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"gorm.io/gorm"
)

type BookRepositoryGorm struct {
	client *gorm.DB
}

// NewBookRepositoryGorm create a new instance of BookRepositoryGorm
func NewBookRepositoryGorm(dbClient *gorm.DB) BookRepositoryGorm {
	return BookRepositoryGorm{dbClient}
}

// SaveBook save book in database
func (r BookRepositoryGorm) SaveBook(book *domain.Book) *errs.AppError {
	if err := r.client.Create(book).Error; err != nil {
		logger.Error(err.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	return nil
}

// FindAllBook find all book in database
func (r BookRepositoryGorm) FindAllBook() ([]domain.Book, *errs.AppError) {
	Books := []domain.Book{}
	if err := r.client.Find(&Books).Error; err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return Books, nil
}

// FindBookById find book by ID in database
func (r BookRepositoryGorm) FindBookById(id uint) (*domain.Book, *errs.AppError) {
	var Book *domain.Book

	if err := r.client.Where("id = ?", id).First(&Book).Error; err != nil {
		logger.Error(err.Error())
		if strings.Contains(err.Error(), "record not found") {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return Book, nil
}

// UpdateBook update book in database
func (r BookRepositoryGorm) UpdateBook(Book *domain.Book) (*domain.Book, *errs.AppError) {
	var result *gorm.DB
	if result = r.client.Where("id = ?", Book.ID).Updates(&Book); result.Error != nil {
		logger.Error(result.Error.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// validates if the rows have changed
	if result.RowsAffected < 1 {
		logger.Info(fmt.Sprintf("Row with id=%d cannot be updated because it doesn't exist", Book.ID))
		return nil, errs.NewNotFoundError("Book not found")
	}

	return Book, nil
}

// DeleteBook delete book in database
func (r BookRepositoryGorm) DeleteBook(id uint) *errs.AppError {
	var Book *domain.Book
	var result *gorm.DB
	if result = r.client.Where("id = ?", id).Delete(&Book); result.Error != nil {
		logger.Error(result.Error.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}

	// validates if the rows have changed
	if result.RowsAffected < 1 {
		logger.Info(fmt.Sprintf("Row with id=%d cannot be deleted because it doesn't exist", id))
		return errs.NewNotFoundError("Book not found")
	}

	return nil
}
