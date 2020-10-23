package stuntman

import (
	entities "crud/src/business/entities"

	"github.com/stretchr/testify/mock"
)

// BookSpy ...
type BookSpy struct {
	mock.Mock
}

// RegisterABook ...
func (u *BookSpy) RegisterABook(book *entities.InputCreateBook) (*entities.BookEntity, error) {

	args := u.Called(book)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// FindABook ...
func (u *BookSpy) FindABook(id uint64) (*entities.BookEntity, error) {

	args := u.Called(id)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// GetAllBooks ...
func (u *BookSpy) GetAllBooks() (*[]entities.BookEntity, error) {

	args := u.Called()

	return args.Get(0).(*[]entities.BookEntity), args.Error(1)
}

// UpdateABook ...
func (u *BookSpy) UpdateABook(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error) {

	args := u.Called(id, book)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// DeleteABook ...
func (u *BookSpy) DeleteABook(id uint64) (*entities.BookEntity, error) {

	args := u.Called(id)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}
