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

type mocksIndexStruct struct {
	bookEntity *entities.BookEntity
	request    *core.HTTPRequest
	usecaseSpy *stuntman.BookSpy
}

func makeIndexMocks() *mocksIndexStruct {

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

	return &mocksIndexStruct{
		bookEntity: bookEntity,
		request:    request,
		usecaseSpy: usecaseSpy,
	}
}

func TestIndexController(t *testing.T) {
	mocks := makeIndexMocks()
	mocks.usecaseSpy.On("FindABook", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := IndexBookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 200)
}

func TestIndecIfErrorOccurWhenTryConverteParamsID(t *testing.T) {
	mocks := makeIndexMocks()
	mocks.usecaseSpy.On("FindABook", mocks.bookEntity.ID).Return(mocks.bookEntity, nil)

	sut := IndexBookController(mocks.usecaseSpy)
	response := sut.Handle(&core.HTTPRequest{})

	assert.Equal(t, response.StatusCode, 500)
}

func TestIndexIfSomeErrorOccurOnUsecaseReturnsStatusCode500(t *testing.T) {
	mocks := makeIndexMocks()
	mocks.usecaseSpy.On("FindABook", mocks.bookEntity.ID).Return(mocks.bookEntity, &errors.InternalServerError{})

	sut := IndexBookController(mocks.usecaseSpy)
	response := sut.Handle(mocks.request)

	assert.Equal(t, response.StatusCode, 500)
}

// func TestIfResultIsNilReturnStatusCode404(t *testing.T) {
// 	mocks := makeIndexMocks()
// 	mocks.usecaseSpy.On("FindABook", mocks.bookEntity.ID).Return(nil, nil)

// 	sut := IndexBookController(mocks.usecaseSpy)
// 	response := sut.Handle(mocks.request)

// 	assert.Equal(t, response.StatusCode, 404)
// }
