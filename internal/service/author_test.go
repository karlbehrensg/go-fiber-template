package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	realDomain "github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/mocks/domain"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
)

var mockAuthorRepo *domain.MockAuthorRepository
var authorService AuthorService

func authorSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAuthorRepo = domain.NewMockAuthorRepository(ctrl)
	authorService = NewAuthorService(mockAuthorRepo)
	return func() {
		authorService = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_author_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	req := requests.AuthorRequest{
		FullName: "J. J. Benitez",
	}

	a := &realDomain.Author{
		FullName: "J. J. Benitez",
	}

	mockAuthorRepo.EXPECT().SaveAuthor(a).Return(errs.NewUnexpectedError("Unexpected database error"))
	// Act
	appError := authorService.CreateAuthor(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new author")
	}

}

func Test_should_return_message_response_when_a_new_author_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	req := requests.AuthorRequest{
		FullName: "J. J. Benitez",
	}

	a := &realDomain.Author{
		FullName: "J. J. Benitez",
	}

	mockAuthorRepo.EXPECT().SaveAuthor(a).Return(nil)
	// Act
	appError := authorService.CreateAuthor(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new author")
	}
}

func Test_should_return_an_error_from_the_server_side_when_get_all_author(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorRepo.EXPECT().FindAllAuthor().Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := authorService.FindAllAuthor()

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for list author")
	}
}

func Test_should_return_successfully_response_when_list_author(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	authors := []realDomain.Author{
		{FullName: "J. J. Benitez"},
		{FullName: "Gabriel García Márquez"},
	}

	mockAuthorRepo.EXPECT().FindAllAuthor().Return(authors, nil)

	// Act
	listAuthor, appError := authorService.FindAllAuthor()

	// Assert
	if appError != nil {
		t.Error("Test failed while get author list")
	}
	if len(listAuthor) == 0 {
		t.Error("Failed while geting author list")
	}
}

func Test_should_return_an_error_from_the_server_when_get_author_by_id(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorRepo.EXPECT().FindAuthorById(uint(2)).Return(nil, errs.NewNotFoundError("record not found"))
	// Act
	_, appError := authorService.FindAuthorById(2)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for author")
	}
}

func Test_should_return_successfully_response_when_find_author_by_id(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	a := &realDomain.Author{
		ID: 2, FullName: "J. J. Benitez",
	}

	mockAuthorRepo.EXPECT().FindAuthorById(uint(2)).Return(a, nil)

	// Act
	author, appError := authorService.FindAuthorById(2)

	// Assert
	if appError != nil {
		t.Error("Test failed while get author")
	}
	if author.Id != a.ID {
		t.Error("Failed while geting author")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_author_cannot_be_updated(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	req := &requests.AuthorRequest{
		FullName: "J. J. Benitez",
	}

	a := &realDomain.Author{
		FullName: "J. J. Benitez",
	}
	mockAuthorRepo.EXPECT().UpdateAuthor(a).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := authorService.UpdateAuthor(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for author")
	}

}

func Test_should_return_author_response_when_a_author_is_updated_successfully(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	req := &requests.AuthorRequest{
		Id:       2,
		FullName: "J. J. Benitez 1",
	}

	a := &realDomain.Author{
		FullName: "J. J. Benitez 1",
	}

	authorWithId := a
	authorWithId.ID = 2
	mockAuthorRepo.EXPECT().UpdateAuthor(a).Return(authorWithId, nil)
	// Act
	updateAuthor, appError := authorService.UpdateAuthor(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while updating author")
	}

	if updateAuthor.FullName != authorWithId.FullName {
		t.Error("Failed while mathching author")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_author_cannot_be_deleted(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorRepo.EXPECT().DeleteAuthor(uint(1)).Return(errs.NewUnexpectedError("Unexpected database error"))
	// Act
	appError := authorService.DeleteAuthor(1)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for author")
	}

}

func Test_should_return_anything_error_response_when_a_author_is_deleted_successfully(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorRepo.EXPECT().DeleteAuthor(uint(1)).Return(nil)
	// Act
	appError := authorService.DeleteAuthor(1)

	// Assert
	if appError != nil {
		t.Error("Test failed while deleting author")
	}
}
