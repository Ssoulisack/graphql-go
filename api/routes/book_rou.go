package routes

import (
	"go-fiber/api/controller"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// BookRoutes sets up book-related GraphQL routes
func BookRoutes(api fiber.Router, db *gorm.DB) {
	br := repositories.NewBookRepository(db)
	bs := services.NewBookService(br)
	bc := controller.NewBookController(bs)

	api.Post("/books", bc.CreateBook)
	api.Get("/books", bc.GetAllBooks)
	api.Get("/books/:id", bc.GetBookByID)
	api.Put("/books/:id", bc.UpdateBook)
	api.Delete("/books/:id", bc.DeleteBook)

}
