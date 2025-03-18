package routes

import (
	"go-fiber/api/graph/resolver"
	"go-fiber/data/repositories"
	"go-fiber/data/services"
	"gorm.io/gorm"
)

// InitializeServices initializes repositories, services, and resolver
func InitializeServices(db *gorm.DB) *resolver.Resolver {
	// Initialize repositories
	authorRepo := repositories.NewAuthorRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	// Initialize services
	authorService := services.NewAuthorService(authorRepo)
	bookService := services.NewBookService(bookRepo)

	// Create resolver instance
	return resolver.NewResolver(authorService, bookService)
}
