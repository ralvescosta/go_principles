package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func createBook(ctx *fiber.Ctx, books *[]Book) error {
	u := new(Book)

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	if u.Author == "" || u.Title == "" {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	lastBook := (*books)[len(*books)-1]

	u.Id = lastBook.Id + 1

	*books = append(*books, *u)

	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "created"}
	return ctx.Status(http.StatusCreated).JSON(response)
}
