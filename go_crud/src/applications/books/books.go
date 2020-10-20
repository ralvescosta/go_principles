package books

import (
	errors "crud/src/__core__/errors"
	entities "crud/src/business/entities"
	repositories "crud/src/frameworks/repositories"
)

type usecase struct {
	repository repositories.IBooksRepository
}

func (u *usecase) RegisterABook(book *entities.InputCreateBook) (*entities.BookEntity, error) {

	result, err := u.repository.Create(book)

	if err != nil {
		return nil, &errors.InternalServerError{}
	}

	return result, nil
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
func Books(repository repositories.IBooksRepository) IBooks {
	return &usecase{repository: repository}
}
