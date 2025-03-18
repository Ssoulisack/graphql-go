package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
)

type BookService interface {
	GetAllBooks() ([]entities.Book, error)
	GetBookByID(id uint) (entities.Book, error)
	GetBooksByAuthorID(authorID uint) ([]entities.Book, error)
	CreateBook(book *entities.Book) error
	UpdateBook(book *entities.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetAllBooks() ([]entities.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id uint) (entities.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) GetBooksByAuthorID(authorID uint) ([]entities.Book, error) {
	return s.repo.FindByAuthorID(authorID)
}

func (s *bookService) CreateBook(book *entities.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *entities.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
