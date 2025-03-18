package controller

import (
	"go-fiber/data/services"

	"github.com/gofiber/fiber/v2"
)

type AuthorController interface {
	CreateAuthor(ctx *fiber.Ctx) error
	GetAllAuthors(ctx *fiber.Ctx) error
	GetAuthorByID(ctx *fiber.Ctx) error
	UpdateAuthor(ctx *fiber.Ctx) error
	DeleteAuthor(ctx *fiber.Ctx) error
}

type authorController struct {
	as services.AuthorService
}

func (ac *authorController) CreateAuthor(ctx *fiber.Ctx) error {
	return nil
}

func (ac *authorController) GetAllAuthors(ctx *fiber.Ctx) error {
	return nil
}

func (ac *authorController) GetAuthorByID(ctx *fiber.Ctx) error {
	return nil
}

func (ac *authorController) UpdateAuthor(ctx *fiber.Ctx) error {
	return nil
}

func (ac *authorController) DeleteAuthor(ctx *fiber.Ctx) error {
	return nil
}

func NewAuthorController(as services.AuthorService) AuthorController {
	return &authorController{as}
}