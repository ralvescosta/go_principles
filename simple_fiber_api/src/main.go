package main

import (
	"simple_fiber_api/src/books"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	booksInMemory := []books.Book{
		books.Book{Id: 1, Author: "fulano", Title: "Livro"},
	}

	books.BooksModule(server, booksInMemory)

	server.Listen(":4000")
}
