package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]entities.Book, error)
	FindByID(id uint) (entities.Book, error)
	FindByAuthorID(authorID uint) ([]entities.Book, error)
	Create(book *entities.Book) error
	Update(book *entities.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]entities.Book, error) {
	var books []entities.Book
	result := r.db.Preload("Author").Find(&books)
	return books, result.Error
}

func (r *bookRepository) FindByID(id uint) (entities.Book, error) {
	var book entities.Book
	result := r.db.Preload("Author").First(&book, id)
	return book, result.Error
}

func (r *bookRepository) FindByAuthorID(authorID uint) ([]entities.Book, error) {
	var books []entities.Book
	result := r.db.Where("author_id = ?", authorID).Find(&books)
	return books, result.Error
}

func (r *bookRepository) Create(book *entities.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *entities.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Book{}, id).Error
}
