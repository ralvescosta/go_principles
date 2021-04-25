package books

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type HttpResponseOnPOST struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

func BooksModule(server *fiber.App, books *[]Book) {

	server.Post("/books", func(ctx *fiber.Ctx) error {
		return createBook(ctx, books)
	})

	server.Get("/books", func(ctx *fiber.Ctx) error {
		return getAllBooks(ctx, books)
	})

	server.Get("/books/:id", func(ctx *fiber.Ctx) error {
		return getBook(ctx, books)
	})

	server.Put("/books/:id", func(ctx *fiber.Ctx) error {
		return updateBook(ctx, books)
	})

	server.Delete("/books/:id", func(ctx *fiber.Ctx) error {
		return deleteBook(ctx, books)
	})
}
