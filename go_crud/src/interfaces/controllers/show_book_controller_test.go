package controllers

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	core "crud/src/__core__"
	errors "crud/src/__core__/errors"
	entities "crud/src/business/entities"
	stuntman "crud/src/interfaces/__test__"
)

type mocksShowStruct struct {
	booksEntity *[]entities.BookEntity
	request     *core.HTTPRequest
	usecaseSpy  *stuntman.BookSpy
}

func makeShowMocks() *mocksShowStruct {

	booksEntity := &[]entities.BookEntity{
		entities.BookEntity{
			ID:                1,
			Title:             "title",
			Author:            "author",
			PublishingCompany: "publishing",
			Edition:           1,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         sql.NullTime{},
		},
	}

	request := &core.HTTPRequest{}

	usecaseSpy := &stuntman.BookSpy{}

	return &mocksShowStruct{
		booksEntity: booksEntity,
		request:     request,
		usecaseSpy:  usecaseSpy,
	}
}

func TestShowController(t *testing.T) {
	mocks := makeShowMocks()
	mocks.usecaseSpy.On("GetAllBooks").Return(mocks.booksEntity, nil)

	sut := ShowBookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 200)
}

func TestShowIfSomeErrorOccurOnUsecaseReturnsStatusCode500(t *testing.T) {
	mocks := makeShowMocks()
	mocks.usecaseSpy.On("GetAllBooks").Return(mocks.booksEntity, &errors.InternalServerError{})

	sut := ShowBookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 500)
}

// func TestIfResultIsNilReturnStatusCode404(t *testing.T) {
// 	mocks := makeShowMocks()
// 	mocks.usecaseSpy.On("FindABook", mocks.bookEntity.ID).Return(nil, nil)

// 	sut := IndexBookController(mocks.usecaseSpy)
// 	response := sut.Handle(mocks.request)

// 	assert.Equal(t, response.StatusCode, 404)
// }
