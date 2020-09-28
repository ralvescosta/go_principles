package books

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func updateBook(ctx *fiber.Ctx, books *[]Book) error {
	strId := ctx.Params("id")
	intId, _ := strconv.ParseInt(strId, 6, 12)
	intId -= 1

	if intId < 0 {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	u := new(Book)

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	if u.Author != "" {
		(*books)[intId].Author = u.Author
	}
	if u.Title != "" {
		(*books)[intId].Title = u.Title
	}

	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "updated"}
	return ctx.JSON(response)
}
