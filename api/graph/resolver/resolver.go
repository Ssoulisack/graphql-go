package resolver

import (
	"go-fiber/data/repositories"
	"go-fiber/data/services"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authorSvc services.AuthorService
	bookSvc   services.BookService
}
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
// NewResolver creates a new Resolver instance
func NewResolver(authorSvc services.AuthorService, bookSvc services.BookService) *Resolver {
	return &Resolver{authorSvc, bookSvc}
}

// InitializeServices initializes repositories, services, and resolver
func InitializeServices(db *gorm.DB) *Resolver {
	// Initialize repositories
	authorRepo := repositories.NewAuthorRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	// Initialize services
	authorService := services.NewAuthorService(authorRepo)
	bookService := services.NewBookService(bookRepo)

	// Create resolver instance
	return NewResolver(authorService, bookService)
}
