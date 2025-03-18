package controller

import (
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	CreateBook(ctx *fiber.Ctx) error
	GetAllBooks(ctx *fiber.Ctx) error
	GetBookByID(ctx *fiber.Ctx) error
	UpdateBook(ctx *fiber.Ctx) error
	DeleteBook(ctx *fiber.Ctx) error
}

type bookController struct {
	bs services.BookService
}

func (bc *bookController) CreateBook(ctx *fiber.Ctx) error {
	return nil
}

func (bc *bookController) GetAllBooks(ctx *fiber.Ctx) error {
	return nil
}

func (bc *bookController) GetBookByID(ctx *fiber.Ctx) error {
	return nil
}

func (bc *bookController) UpdateBook(ctx *fiber.Ctx) error {
	return nil
}

func (bc *bookController) DeleteBook(ctx *fiber.Ctx) error {
	return nil
}

func NewBookController(bs services.BookService) BookController {
	return &bookController{bs}
}
