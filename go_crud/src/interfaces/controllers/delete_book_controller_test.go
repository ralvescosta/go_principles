package controllers

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	stuntman "crud/src/interfaces/__test__"

	core "crud/src/__core__"
	errors "crud/src/__core__/errors"
	entities "crud/src/business/entities"
)

type mocksDeleteStruct struct {
	inputCreateBook *entities.InputCreateBook
	bookEntity      *entities.BookEntity
	request         *core.HTTPRequest
	usecaseSpy      *stuntman.BookSpy
}

func makeDeleteMocks() *mocksDeleteStruct {

	inputCreateBook := &entities.InputCreateBook{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "publishing",
		Edition:           1,
	}

	bookEntity := &entities.BookEntity{
		ID:                1,
		Title:             "title",
		Author:            "author",
		PublishingCompany: "publishing",
		Edition:           1,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         sql.NullTime{},
	}

	param := make(map[string]string)
	param["id"] = "1"
	request := &core.HTTPRequest{
		Params: param,
	}

	usecaseSpy := &stuntman.BookSpy{}

	return &mocksDeleteStruct{
		inputCreateBook: inputCreateBook,
		bookEntity:      bookEntity,
		request:         request,
		usecaseSpy:      usecaseSpy,
	}
}

func TestDeleteController(t *testing.T) {
	mocks := makeDeleteMocks()
	mocks.usecaseSpy.On("DeleteABook", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := DeleteABookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 200)
}

func TestDeleteIfErrorOccurWhenTryConverteParamsID(t *testing.T) {
	mocks := makeDeleteMocks()
	mocks.usecaseSpy.On("DeleteABook", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := DeleteABookController(mocks.usecaseSpy)
	response := sut.Handle(&core.HTTPRequest{})

	assert.Equal(t, response.StatusCode, 500)
}

func TestDeleteIfSomeErrorOccurOnUsecaseReturnsStatusCode500(t *testing.T) {
	mocks := makeDeleteMocks()
	mocks.usecaseSpy.On("DeleteABook", mocks.bookEntity.ID).Return(mocks.bookEntity, &errors.InternalServerError{})

	sut := DeleteABookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 500)
}
