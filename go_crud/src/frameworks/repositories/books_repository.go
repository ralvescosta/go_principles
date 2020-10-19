package repositories

import (
	"database/sql"
	"fmt"

	entities "crud/src/business/entities"
)

type booksRepository struct {
	db *sql.DB
}

// Create ...
func (b *booksRepository) Create(book *entities.InputCreateBook) (*entities.BookEntity, error) {
	sql := `INSERT INTO "books" (title, author, publishing_company, edition) VALUES ($1, $2, $3, $4) RETURNING`

	prepare, err := (*b.db).Prepare(sql)
	if err != nil {
		return nil, err
	}

	result, err := prepare.Exec(book.Title, book.Author, book.PublishingCompany, book.Edition)
	if err != nil {
		return nil, err
	}

	fmt.Println(result)

	return nil, nil
}

// FindByID ...
func (b *booksRepository) FindByID(id uint64) (*entities.BookEntity, error) {
	return nil, nil
}

// FindAll ...
func (b *booksRepository) FindAll() (*[]entities.BookEntity, error) {
	return nil, nil
}

// Update ...
func (b *booksRepository) Update(*entities.BookEntity) (*entities.BookEntity, error) {
	return nil, nil
}

// Delete ...
func (b *booksRepository) Delete(id uint64) (*entities.BookEntity, error) {
	return nil, nil
}

// BooksRepository ...
func BooksRepository() IBooksRepository {
	return &booksRepository{}
}
