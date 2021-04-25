package books

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func deleteBook(ctx *fiber.Ctx, books *[]Book) error {
	strId := ctx.Params("id")
	intId, _ := strconv.ParseInt(strId, 6, 12)

	if intId < 0 {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	if (*books)[intId-1].Id != intId {
		return ctx.Status(http.StatusBadRequest).SendString("")
	}

	*books = append((*books)[:intId-1], (*books)[intId:]...)

	response := HttpResponseOnPOST{StatusCode: http.StatusCreated, Message: "deleted"}
	return ctx.JSON(response)
}
