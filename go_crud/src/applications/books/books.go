package books

type usecase struct{}

func (*usecase) RegisterABook() {}
func (*usecase) FindABook()     {}
func (*usecase) GetAllBooks()   {}
func (*usecase) UpdateABook()   {}
func (*usecase) DeleteABook()   {}

// Books ...
func Books() IBooks {
	return &usecase{}
}
