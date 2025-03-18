package routes

import (
	"go-fiber/api/controller"
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AuthorRoutes sets up author-related GraphQL routes
func AuthorRoutes(api fiber.Router, db *gorm.DB) {
	ar := repositories.NewAuthorRepository(db)
	as := services.NewAuthorService(ar)
	ac := controller.NewAuthorController(as)

	api.Get("/authors", ac.GetAllAuthors)
	api.Get("/author/:id", ac.GetAuthorByID)
	api.Post("/author", ac.CreateAuthor)
	api.Put("/author/:id", ac.UpdateAuthor)
	api.Delete("/author/:id", ac.DeleteAuthor)

}
