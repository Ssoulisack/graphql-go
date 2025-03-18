package models

import "go-fiber/domain/entities"

type BookRequest struct {
	Title    string `json:"title" binding:"required"`
	AuthorID uint   `json:"author_id" binding:"required"`
}

type BookResponse struct {
	ID       uint           `json:"id"`
	Title    string         `json:"title"`
	AuthorID uint           `json:"author_id"`
	Author   AuthorResponse `json:"author,omitempty"`
}

func BookToResponse(book entities.Book) BookResponse {
	return BookResponse{
		ID:       book.ID,
		Title:    book.Title,
		AuthorID: book.AuthorID,
		Author:   AuthorToResponse(book.Author),
	}
}
