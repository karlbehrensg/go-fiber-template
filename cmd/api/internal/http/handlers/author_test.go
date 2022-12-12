package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/karlbehrensg/go-fiber-template/internal/http/requests"
	"github.com/karlbehrensg/go-fiber-template/internal/http/responses"
	"github.com/karlbehrensg/go-fiber-template/mocks/service"
	"github.com/karlbehrensg/go-fiber-template/pkg/errs"
	"github.com/stretchr/testify/assert"
)

var router *fiber.App
var ah AuthorHandler
var mockAuthorService *service.MockAuthorService

func authorSetup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAuthorService = service.NewMockAuthorService(ctrl)
	ah = AuthorHandler{Service: mockAuthorService}
	router = fiber.New()
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_author_with_status_code_201(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	author := requests.AuthorRequest{
		FullName: "J. J. Benitez",
	}

	authorBytes := new(bytes.Buffer)
	json.NewEncoder(authorBytes).Encode(author)

	mockAuthorService.EXPECT().CreateAuthor(author).Return(nil)
	router.Post("/author", ah.CreateAuthor)
	request, _ := http.NewRequest(http.MethodPost, "/author", authorBytes)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Author created", data.Message, "The message should be the same.")

}

func Test_should_return_status_code_400_with_error_message_when_call_CreateAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	rawPayload := []byte(`{"full_name":}`)

	router.Post("/author", ah.CreateAuthor)
	request, _ := http.NewRequest(
		http.MethodPost,
		"/author",
		bytes.NewReader(rawPayload),
	)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Invalid data", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_400_with_error_message_when_validate_struct_in_CreateAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	rawPayload := []byte(`{"full_name": "ed"}`)

	router.Post("/author", ah.CreateAuthor)
	request, _ := http.NewRequest(
		http.MethodPost,
		"/author",
		bytes.NewReader(rawPayload),
	)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(
		t,
		"Key: 'AuthorRequest.FullName' Error:Field validation for 'FullName' failed on the 'min' tag",
		data.Message,
		"The message should be the same.",
	)
}

func Test_should_return_status_code_500_with_error_message_when_call_CreateAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	author := requests.AuthorRequest{
		FullName: "J. J. Benitez",
	}

	rawPayload := []byte(`{"full_name": "J. J. Benitez"}`)

	mockAuthorService.EXPECT().CreateAuthor(author).Return(errs.NewUnexpectedError("Unexpected error from database"))
	router.Post("/author", ah.CreateAuthor)
	request, _ := http.NewRequest(
		http.MethodPost,
		"/author",
		bytes.NewReader(rawPayload),
	)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Test failed with different status code")
	assert.Equal(
		t,
		"Unexpected error from database",
		data.Message,
		"The message should be the same.",
	)
}

func Test_should_return_status_code_500_with_error_message_when_call_GetAllAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorService.EXPECT().FindAllAuthor().Return(nil, errs.NewUnexpectedError("Unexpected error from database"))
	router.Get("/author", ah.GetAllAuthor)
	request, _ := http.NewRequest(http.MethodGet, "/author", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Unexpected error from database", data.Message, "The message should be the same.")
}

func Test_should_return_authors_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	authors := []responses.AuthorResponse{
		{Id: 1, FullName: "J. J. Benitez"},
		{Id: 2, FullName: "Gabriel Garcia Marquez"},
	}

	mockAuthorService.EXPECT().FindAllAuthor().Return(authors, nil)
	router.Get("/author", ah.GetAllAuthor)
	request, _ := http.NewRequest(http.MethodGet, "/author", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data []responses.AuthorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Test failed with different status code")
	assert.Len(t, authors, 2, "Failed test Len is less than 2")
}

func Test_should_return_author_with_status_code_200_when_call_GetAuthorById(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	author := responses.AuthorResponse{
		Id:       1,
		FullName: "J. J. Benitez",
	}

	mockAuthorService.EXPECT().FindAuthorById(uint(1)).Return(&author, nil)
	router.Get("/author/:id", ah.GetAuthorById)
	request, _ := http.NewRequest(http.MethodGet, "/author/1", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.AuthorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, author, data, "The message should be the same.")
}

func Test_should_return_status_code_400_when_call_GetAuthorById_with_bad_id(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	router.Get("/author/:id", ah.GetAuthorById)
	request, _ := http.NewRequest(http.MethodGet, "/author/el", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Invalid author id", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_500_when_call_GetAuthorById(t *testing.T) {
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorService.EXPECT().FindAuthorById(uint(1)).Return(nil, errs.NewNotFoundError("record not found"))
	router.Get("/author/:id", ah.GetAuthorById)
	request, _ := http.NewRequest(http.MethodGet, "/author/1", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "record not found", data.Message, "The message should be the same.")
}

func Test_should_return_author_with_status_code_200_when_updated(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	author := requests.AuthorRequest{
		Id:       1,
		FullName: "J. J. Benitez",
	}

	authorResp := responses.AuthorResponse{
		Id:       1,
		FullName: "J. J. Benitez",
	}

	authorBytes := new(bytes.Buffer)
	json.NewEncoder(authorBytes).Encode(author)

	mockAuthorService.EXPECT().UpdateAuthor(&author).Return(&authorResp, nil)
	router.Put("/author/:id", ah.UpdateAuthor)
	request, _ := http.NewRequest(http.MethodPut, "/author/1", authorBytes)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.AuthorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, authorResp, data, "The message should be the same.")

}

func Test_should_return_status_code_400_when_call_UpdateAuthor_with_bad_id(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	router.Put("/author/:id", ah.UpdateAuthor)
	request, _ := http.NewRequest(http.MethodPut, "/author/a", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Invalid author id", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_400_when_validate_struct_in_UpdateAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	router.Put("/author/:id", ah.UpdateAuthor)
	request, _ := http.NewRequest(http.MethodPut, "/author/1", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Invalid data", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_400_when_call_UpdateAuthor_with_bad_request(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	rawPayload := []byte(`{"full_name": "ed"}`)

	router.Put("/author/:id", ah.UpdateAuthor)
	request, _ := http.NewRequest(
		http.MethodPut,
		"/author/1",
		bytes.NewReader(rawPayload),
	)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Key: 'AuthorRequest.FullName' Error:Field validation for 'FullName' failed on the 'min' tag", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_500_when_call_UpdateAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	author := requests.AuthorRequest{
		Id:       1,
		FullName: "J. J. Benitez",
	}

	rawPayload := []byte(`{"id": 1, "full_name": "J. J. Benitez"}`)

	authorBytes := new(bytes.Buffer)
	json.NewEncoder(authorBytes).Encode(author)

	mockAuthorService.EXPECT().UpdateAuthor(&author).Return(nil, errs.NewUnexpectedError("key full_name duplicate value"))
	router.Put("/author/:id", ah.UpdateAuthor)
	request, _ := http.NewRequest(
		http.MethodPut,
		"/author/1",
		bytes.NewReader(rawPayload),
	)
	request.Header.Add("Content-Type", "application/json")

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "key full_name duplicate value", data.Message, "The message should be the same.")
}

func Test_should_deleted_author_and_return_status_code_204(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorService.EXPECT().DeleteAuthor(uint(1)).Return(nil)
	router.Delete("/author/:id", ah.DeleteAuthor)
	request, _ := http.NewRequest(http.MethodDelete, "/author/1", nil)

	// Act
	resp, _ := router.Test(request, -1)
	// Assert
	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "Test failed with different status code")

}

func Test_should_return_status_code_400_when_call_DeleteAuthor_with_bad_id(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	router.Delete("/author/:id", ah.DeleteAuthor)
	request, _ := http.NewRequest(http.MethodDelete, "/author/a", nil)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Invalid author id", data.Message, "The message should be the same.")
}

func Test_should_return_status_code_500_when_call_DeleteAuthor(t *testing.T) {
	// Arrange
	teardown := authorSetup(t)
	defer teardown()

	mockAuthorService.EXPECT().DeleteAuthor(uint(1)).Return(errs.NewUnexpectedError("Unexpected error from database"))
	router.Put("/author/:id", ah.DeleteAuthor)
	request, _ := http.NewRequest(
		http.MethodPut,
		"/author/1",
		nil,
	)

	// Act
	resp, _ := router.Test(request, -1)
	decoder := json.NewDecoder(resp.Body)
	var data responses.ErrorResponse
	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Test failed while ummarshal resp.Body")
	}

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode, "Test failed with different status code")
	assert.Equal(t, "Unexpected error from database", data.Message, "The message should be the same.")
}
