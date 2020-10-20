package books

import (
	entities "crud/src/business/entities"
)

type usecase struct{}

func (*usecase) RegisterABook(book *entities.InputCreateBook) (*entities.BookEntity, error) {
	return nil, nil
}

func (*usecase) FindABook(id uint64) (*entities.BookEntity, error) {
	return nil, nil
}

func (*usecase) GetAllBooks() (*[]entities.BookEntity, error) {
	return nil, nil
}

func (*usecase) UpdateABook(book *entities.BookEntity) (*entities.BookEntity, error) {
	return nil, nil
}

func (*usecase) DeleteABook(id uint64) (*entities.BookEntity, error) {
	return nil, nil
}

// Books ...
func Books() IBooks {
	return &usecase{}
}
