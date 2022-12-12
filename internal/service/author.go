package service

import (
	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
)

// AuthorService port primary
//
//go:generate mockgen -destination=../../mocks/service/mockAuthorService.go -package=service github.com/karlbehrensg/go-fiber-template/internal/service AuthorService
type AuthorService interface {
	CreateAuthor(requests.AuthorRequest) *errs.AppError
	FindAllAuthor() ([]responses.AuthorResponse, *errs.AppError)
	FindAuthorById(uint) (*responses.AuthorResponse, *errs.AppError)
	UpdateAuthor(*requests.AuthorRequest) (*responses.AuthorResponse, *errs.AppError)
	DeleteAuthor(uint) *errs.AppError
}

type DefaultAuthorService struct {
	repo domain.AuthorRepository
}

// NewAuthorService create a new instance of DefaultAuthorService
func NewAuthorService(repository domain.AuthorRepository) DefaultAuthorService {
	return DefaultAuthorService{repository}
}

// CreateAuthor use case for create author
func (s DefaultAuthorService) CreateAuthor(request requests.AuthorRequest) *errs.AppError {

	author := &domain.Author{
		FullName: request.FullName,
	}

	// calls repository to save author
	if err := s.repo.SaveAuthor(author); err != nil {
		return err
	}

	return nil
}

// FindAllAuthor use case for find all author
func (s DefaultAuthorService) FindAllAuthor() ([]responses.AuthorResponse, *errs.AppError) {
	var authors []domain.Author
	var err *errs.AppError
	// calls repository to find all author
	if authors, err = s.repo.FindAllAuthor(); err != nil {
		return nil, err
	}

	response := make([]responses.AuthorResponse, 0)
	for _, author := range authors {
		response = append(response, *author.ToNewAuthorResponse())
	}

	return response, nil
}

// FindAuthorById use case for find author by ID
func (s DefaultAuthorService) FindAuthorById(id uint) (*responses.AuthorResponse, *errs.AppError) {
	var author *domain.Author
	var err *errs.AppError
	// calls repository to find author by ID
	if author, err = s.repo.FindAuthorById(id); err != nil {
		return nil, err
	}

	response := *author.ToNewAuthorResponse()

	return &response, nil
}

// UpdateAuthor use case for update author
func (s DefaultAuthorService) UpdateAuthor(request *requests.AuthorRequest) (*responses.AuthorResponse, *errs.AppError) {
	author := &domain.Author{
		ID:       request.Id,
		FullName: request.FullName,
	}

	var err *errs.AppError
	// calls repository to update author
	if author, err = s.repo.UpdateAuthor(author); err != nil {
		return nil, err
	}

	response := *author.ToNewAuthorResponse()

	return &response, nil

}

// DeleteAuthor use case for delete author
func (s DefaultAuthorService) DeleteAuthor(id uint) *errs.AppError {
	// calls repository to delete author
	if err := s.repo.DeleteAuthor(id); err != nil {
		return err
	}

	return nil

}
