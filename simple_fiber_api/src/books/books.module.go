package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Id     int16  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type HttpResponseOnPOST struct {
	StatusCode int32  `json:"statusCode"`
	Message    string `json:"message"`
}

type HttpResponseOnGet struct {
	Book []Book `json:"book"`
}

func createBook(ctx *fiber.Ctx, books []Book) error {

	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "created"}

	// append(books, Book{Id: 2, Author: "fulanso", Title: "Livroa"})

	return ctx.Status(http.StatusCreated).JSON(response)
}

func getAllBooks(ctx *fiber.Ctx, books []Book) error {
	return ctx.JSON(books)
}

func getBook(ctx *fiber.Ctx) error {
	response := []Book{
		Book{Id: 1, Author: "fulano", Title: "Livro"},
	}

	return ctx.JSON(response)
}

func updateBook(ctx *fiber.Ctx) error {
	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "created"}
	return ctx.JSON(response)
}

func deleteBook(ctx *fiber.Ctx) error {
	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "created"}
	return ctx.JSON(response)
}

func BooksModule(server *fiber.App, books []Book) {
	server.Post("/books", func(ctx *fiber.Ctx) error {
		return createBook(ctx, books)
	})

	server.Get("/books", func(ctx *fiber.Ctx) error {
		return getAllBooks(ctx, books)
	})

	server.Get("/books/:id", getBook)
	server.Put("/books/:id", updateBook)
	server.Delete("/books/:id", deleteBook)
}
