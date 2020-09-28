package books

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	id     int16
	title  string
	author string
}

type HttpResponse struct {
	statusCode int16
	message    string
}

func createBook(ctx *fiber.Ctx) {
	response := HttpResponse{statusCode: http.StatusCreated, message: "created"}
	ctx.Status(http.StatusCreated).JSON(response)
}

func BooksModule(server *fiber.App) {

	server.Get("/books", func(ctx *fiber.Ctx) error {
		response := HttpResponse{statusCode: http.StatusCreated, message: "created"}
		ctx.JSON(response)
		return nil
	})
	// bookGroup.Get("/", getAllBooks)
	// bookGroup.Get("/:id", getBook)
	// bookGroup.Put("/", updateBook)
	// bookGroup.Delete("/:id", deleteBook)
	fmt.Println("Books LALALALA")
}
