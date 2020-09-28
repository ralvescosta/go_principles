package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HttpResponse struct {
	statusCode int32
	message    string
}

func main() {
	server := fiber.New()

	server.Get("/books", func(ctx *fiber.Ctx) error {
		response := HttpResponse{statusCode: http.StatusCreated, message: "created"}
		return ctx.JSON(response)
	})

	server.Listen(":4000")
}
