package books

import (
	entities "crud/src/business/entities"
)

// IBooks ...
type IBooks interface {
	RegisterABook(book *entities.InputCreateBook)
	FindABook(id uint64)
	GetAllBooks()
	UpdateABook(book *entities.BookEntity)
	DeleteABook(id uint64)
}
