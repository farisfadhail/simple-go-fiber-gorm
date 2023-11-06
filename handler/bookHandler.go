package handler

import (
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)


func BookHandlerGetAll(ctx *fiber.Ctx) error {
	return ctx.SendString("udin")
}

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	err := ctx.BodyParser(book)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(book)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	filename := ctx.Locals("filename").(string)
	if filename == "" {
		return ctx.Status(422).JSON(fiber.Map{
			"message" : "IMAGE COVER IS REQUIRED",
		})
	}

	newBook := entity.Book{
		Title: book.Title,
		Author: book.Author,
		Cover: filename,
	}

	result := db.Debug().Create(&newBook)

	if result.Error != nil {
		utils.HandleRemoveFile(filename)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "FAILED TO STORE DATA",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS CREATE DATA",
		"data" : newBook,
	})
}