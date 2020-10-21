package books

import (
	entities "crud/src/business/entities"
)

// IBooks ...
type IBooks interface {
	RegisterABook(book *entities.InputCreateBook) (*entities.BookEntity, error)
	FindABook(id uint64) (*entities.BookEntity, error)
	GetAllBooks() (*[]entities.BookEntity, error)
	UpdateABook(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error)
	DeleteABook(id uint64) (*entities.BookEntity, error)
}
