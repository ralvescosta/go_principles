package books

import "github.com/gofiber/fiber/v2"

func getAllBooks(ctx *fiber.Ctx, books *[]Book) error {

	return ctx.JSON(books)
}
