package repositories

import entities "crud/src/business/entities"

// IBooksRepository ...
type IBooksRepository interface {
	Create(book *entities.InputCreateBook) (*entities.BookEntity, error)
	FindByID(id uint64) (*entities.BookEntity, error)
	FindAll() (*[]entities.BookEntity, error)
	Update(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error)
	Delete(id uint64) (*entities.BookEntity, error)
	SoftDelete(id uint64) (*entities.BookEntity, error)
}
