package domain

import (
	"time"

	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"gorm.io/gorm"
)

type Book struct {
	ID              uint   `gorm:"id;primary_key"`
	Title           string `gorm:"title;not null"`
	PublicationYear string `gorm:"publication_year"`
	AuthorID        uint   `gorm:"author_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// BookRepository port secondary
type BookRepository interface {
	SaveBook(*Book) *errs.AppError
	FindAllBook() ([]Book, *errs.AppError)
	FindBookById(uint) (*Book, *errs.AppError)
	UpdateBook(*Book) (*Book, *errs.AppError)
	DeleteBook(id uint) *errs.AppError
}

// ToNewBookResponse convert Book struct to responses.BookResponse struct
func (d *Book) ToNewBookResponse() *responses.BookResponse {

	return &responses.BookResponse{
		Id:              d.ID,
		Title:           d.Title,
		PublicationYear: d.PublicationYear,
		AuthorID:        d.AuthorID,
	}
}
