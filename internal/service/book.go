package service

import (
	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
)

// BookService port primary
type BookService interface {
	CreateBook(requests.BookRequest) *errs.AppError
	FindAllBook() ([]responses.BookResponse, *errs.AppError)
	FindBookById(uint) (*responses.BookResponse, *errs.AppError)
	UpdateBook(*requests.BookRequest) (*responses.BookResponse, *errs.AppError)
	DeleteBook(uint) *errs.AppError
}

type DefaultBookService struct {
	repo domain.BookRepository
}

// NewBookService create a new instance of DefaultBookService
func NewBookService(repository domain.BookRepository) DefaultBookService {
	return DefaultBookService{repository}
}

// CreateBook use case for create book
func (s DefaultBookService) CreateBook(request requests.BookRequest) *errs.AppError {
	Book := &domain.Book{
		Title:           request.Title,
		AuthorID:        request.AuthorID,
		PublicationYear: request.PublicationYear,
	}

	// calls repository to save book
	if err := s.repo.SaveBook(Book); err != nil {
		return err
	}

	return nil
}

// FindAllBook use case for find all book
func (s DefaultBookService) FindAllBook() ([]responses.BookResponse, *errs.AppError) {
	var Books []domain.Book
	var err *errs.AppError
	// calls repository to find all book
	if Books, err = s.repo.FindAllBook(); err != nil {
		return nil, err
	}

	response := make([]responses.BookResponse, 0)
	for _, Book := range Books {
		response = append(response, *Book.ToNewBookResponse())
	}

	return response, nil
}

// FindBookById use case for find book by ID
func (s DefaultBookService) FindBookById(id uint) (*responses.BookResponse, *errs.AppError) {
	var Book *domain.Book
	var err *errs.AppError
	// calls repository to find book by ID
	if Book, err = s.repo.FindBookById(id); err != nil {
		return nil, err
	}

	response := *Book.ToNewBookResponse()

	return &response, nil
}

// UpdateAuthor use case for update book
func (s DefaultBookService) UpdateBook(request *requests.BookRequest) (*responses.BookResponse, *errs.AppError) {
	Book := &domain.Book{
		ID:              request.Id,
		Title:           request.Title,
		AuthorID:        request.AuthorID,
		PublicationYear: request.PublicationYear,
	}

	var err *errs.AppError
	// calls repository to update author
	if Book, err = s.repo.UpdateBook(Book); err != nil {
		return nil, err
	}

	response := *Book.ToNewBookResponse()

	return &response, nil

}

// DeleteBook use case for delete book
func (s DefaultBookService) DeleteBook(id uint) *errs.AppError {
	// calls repository to delete book
	if err := s.repo.DeleteBook(id); err != nil {
		return err
	}

	return nil

}
