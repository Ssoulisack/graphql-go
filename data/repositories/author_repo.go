package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]entities.Author, error)
	FindByID(id uint) (entities.Author, error)
	Create(author *entities.Author) error
	Update(author *entities.Author) error
	Delete(id uint) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll() ([]entities.Author, error) {
	var authors []entities.Author
	result := r.db.Find(&authors)
	return authors, result.Error
}

func (r *authorRepository) FindByID(id uint) (entities.Author, error) {
	var author entities.Author
	result := r.db.First(&author, id)
	return author, result.Error
}

func (r *authorRepository) Create(author *entities.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) Update(author *entities.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Author{}, id).Error
}
