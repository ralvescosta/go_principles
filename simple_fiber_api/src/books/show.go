package books

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBook(ctx *fiber.Ctx, books *[]Book) error {
	strId := ctx.Params("id")
	intId, _ := strconv.ParseInt(strId, 6, 12)
	intId -= 1

	if intId < 0 {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	book := (*books)[intId]

	return ctx.JSON(book)
}
