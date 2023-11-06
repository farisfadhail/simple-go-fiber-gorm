package handler

import (
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ImageGalleriesHandlerCreate(ctx *fiber.Ctx) error {
	imageGallery := new(request.ImageCreateRequest)
	err := ctx.BodyParser(imageGallery)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(imageGallery)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	filenames := ctx.Locals("filenames").([]string)
	if len(filenames) == 0 {
		return ctx.Status(422).JSON(fiber.Map{
			"message" : "IMAGE COVER IS REQUIRED",
		})
	}

	for _, filename := range filenames {
		newImageGallery := entity.ImageGallery{
			Image: filename,
			CategoryId: imageGallery.CategoryId,
		}
	
		result := db.Debug().Create(&newImageGallery)
	
		if result.Error != nil {
			log.Println("Some data failed to create")
			return ctx.Status(500).JSON(fiber.Map{
				"message": "FAILED TO STORE DATA",
			})
		}
	}


	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS CREATE DATA",
	})
}

func ImageGalleriesHandlerDestroy(ctx *fiber.Ctx) error {
	imageId := ctx.Params("imageId")

	var image entity.ImageGallery

	result := db.First(&image, imageId)

	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "DATA NOT FOUND",
		})
	}

	// DELETE IMAGE GALLERY DIRECTORY
	err := utils.HandleRemoveFile(image.Image)
	if err != nil {
		log.Println("FAILED TO DELETE FILE IN DIRECTORY")
	}

	// DELETE IMAGE GALLERY DATABASE
	result = db.Debug().Delete(&image, imageId)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "FAILED TO DELETE DATA",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS DELETE DATA",
	})
}