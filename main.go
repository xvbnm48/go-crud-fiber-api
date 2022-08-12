package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xvbnm48/go-crud-fiber-api/controllers/bookControllers"
	"github.com/xvbnm48/go-crud-fiber-api/models"
)

func main() {
	models.ConnectionDatabase()
	app := fiber.New()
	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookControllers.Index)
	book.Get("/:id", bookControllers.Show)
	book.Post("/", bookControllers.Create)
	book.Put("/:id", bookControllers.Update)
	book.Delete("/:id", bookControllers.Delete)

	log.Fatal(app.Listen(":8080"))
}
