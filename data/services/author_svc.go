package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
)

type AuthorService interface {
	GetAllAuthors() ([]entities.Author, error)
	GetAuthorByID(id uint) (entities.Author, error)
	CreateAuthor(author *entities.Author) error
	UpdateAuthor(author *entities.Author) error
	DeleteAuthor(id uint) error
}

type authorService struct {
	repo repositories.AuthorRepository
}

func NewAuthorService(repo repositories.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) GetAllAuthors() ([]entities.Author, error) {
	return s.repo.FindAll()
}

func (s *authorService) GetAuthorByID(id uint) (entities.Author, error) {
	return s.repo.FindByID(id)
}

func (s *authorService) CreateAuthor(author *entities.Author) error {
	return s.repo.Create(author)
}

func (s *authorService) UpdateAuthor(author *entities.Author) error {
	return s.repo.Update(author)
}

func (s *authorService) DeleteAuthor(id uint) error {
	return s.repo.Delete(id)
}
