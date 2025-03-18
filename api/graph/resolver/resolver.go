package resolver

import "go-fiber/data/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authorSvc services.AuthorService
	bookSvc   services.BookService
}

func NewResolver(authorSvc services.AuthorService, bookSvc services.BookService) *Resolver {
	return &Resolver{authorSvc, bookSvc}
}