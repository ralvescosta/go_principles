package books

import entities "crud/src/business/entities"

type usecase struct{}

func (*usecase) RegisterABook(book *entities.InputCreateBook) {}
func (*usecase) FindABook(id uint64)                          {}
func (*usecase) GetAllBooks()                                 {}
func (*usecase) UpdateABook(book *entities.BookEntity)        {}
func (*usecase) DeleteABook(id uint64)                        {}

// Books ...
func Books() IBooks {
	return &usecase{}
}
