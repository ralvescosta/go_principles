package stuntman

import (
	"github.com/stretchr/testify/mock"

	entities "crud/src/business/entities"
)

// BooksRepositorySpy ...
type BooksRepositorySpy struct {
	mock.Mock
}

// Create ...
func (b *BooksRepositorySpy) Create(book *entities.InputCreateBook) (*entities.BookEntity, error) {
	args := b.Called(book)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// FindByID ...
func (b *BooksRepositorySpy) FindByID(id uint64) (*entities.BookEntity, error) {
	args := b.Called(id)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// FindAll ...
func (b *BooksRepositorySpy) FindAll() (*[]entities.BookEntity, error) {
	args := b.Called()

	return args.Get(0).(*[]entities.BookEntity), args.Error(1)
}

// Update ...
func (b *BooksRepositorySpy) Update(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error) {
	args := b.Called(id, book)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// Delete ...
func (b *BooksRepositorySpy) Delete(id uint64) (*entities.BookEntity, error) {
	args := b.Called(id)

	return args.Get(0).(*entities.BookEntity), args.Error(1)
}

// SoftDelete ...
func (b *BooksRepositorySpy) SoftDelete(id uint64) (*entities.BookEntity, error) {

	return nil, nil
}
