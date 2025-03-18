package models

import "go-fiber/domain/entities"

type AuthorRequest struct {
	Name string `json:"name" binding:"required"`
}

type AuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AuthorDetailResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Books []BookResponse `json:"books"`
}

func AuthorToResponse(author entities.Author) AuthorResponse {
	return AuthorResponse{
		ID:   author.ID,
		Name: author.Name,
	}
}

func AuthorToDetailResponse(author entities.Author) AuthorDetailResponse {
	bookResponses := make([]BookResponse, len(author.Books))
	for i, book := range author.Books {
		bookResponses[i] = BookToResponse(book)
	}

	return AuthorDetailResponse{
		ID:    author.ID,
		Name:  author.Name,
		Books: bookResponses,
	}
}
