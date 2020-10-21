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

func (u *usecase) FindABook(id uint64) (*entities.BookEntity, error) {

	result, err := u.repository.FindByID(id)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	if err != nil {
		return nil, &errors.InternalServerError{}
	}

	return result, nil
}

func (u *usecase) GetAllBooks() (*[]entities.BookEntity, error) {

	result, err := u.repository.FindAll()

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	if err != nil {
		return nil, &errors.InternalServerError{}
	}

	return result, nil
}

func (u *usecase) UpdateABook(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error) {

	result, err := u.repository.Update(id, book)
	if err != nil {
		return nil, &errors.InternalServerError{}
	}

	return result, nil
}

func (u *usecase) DeleteABook(id uint64) (*entities.BookEntity, error) {

	result, err := u.repository.Delete(id)
	if err != nil {
		return nil, &errors.InternalServerError{}
	}

	return result, nil
}

// Books ...
func Books(repository repositories.IBooksRepository) IBooks {
	return &usecase{repository: repository}
}
