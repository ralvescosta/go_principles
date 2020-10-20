package repositories

import (
	"database/sql"

	entities "crud/src/business/entities"
)

type booksRepository struct {
	db *sql.DB
}

// Create ...
func (b *booksRepository) Create(book *entities.InputCreateBook) (*entities.BookEntity, error) {
	sql := `INSERT INTO books (title, author, publishing_company, edition) VALUES ($1, $2, $3, $4) RETURNING *`

	prepare, err := (*b.db).Prepare(sql)
	if err != nil {
		return nil, err
	}

	entity := &entities.BookEntity{}
	err = prepare.QueryRow(
		book.Title,
		book.Author,
		book.PublishingCompany,
		book.Edition,
	).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Author,
		&entity.PublishingCompany,
		&entity.Edition,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
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
func BooksRepository(db *sql.DB) IBooksRepository {
	return &booksRepository{db: db}
}
