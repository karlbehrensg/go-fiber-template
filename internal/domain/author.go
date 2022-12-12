package domain

import (
	"time"

	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"gorm.io/gorm"
)

type Author struct {
	ID        uint   `gorm:"id;primary_key"`
	FullName  string `gorm:"full_name;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Books     []Book
}

// AuthorRepository port secondary
//
//go:generate mockgen -destination=../../mocks/domain/mockAuthorRepository.go -package=domain github.com/karlbehrensg/go-fiber-template/internal/domain AuthorRepository
type AuthorRepository interface {
	SaveAuthor(*Author) *errs.AppError
	FindAllAuthor() ([]Author, *errs.AppError)
	FindAuthorById(uint) (*Author, *errs.AppError)
	UpdateAuthor(*Author) (*Author, *errs.AppError)
	DeleteAuthor(id uint) *errs.AppError
}

// ToNewAuthorResponse convert Author struct to responses.AuthorResponse struct
func (d *Author) ToNewAuthorResponse() *responses.AuthorResponse {

	return &responses.AuthorResponse{
		Id:       d.ID,
		FullName: d.FullName,
	}
}
