package books

// IBooks ...
type IBooks interface {
	RegisterABook()
	FindABook()
	GetAllBooks()
	UpdateABook()
	DeleteABook()
}
