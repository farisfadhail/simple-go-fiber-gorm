package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"go-fiber-gorm/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var db = database.DatabaseInit()
var validate = validator.New()

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := db.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(user)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "INTERNAL SERVER ERROR",
		})
	}

	newUser.Password = hashedPassword

	result := db.Debug().Create(&newUser)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "FAILED TO STORE DATA",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: newUser.ID,
	// 	Name: newUser.Name,
	// 	Email: newUser.Email,
	// 	Address: newUser.Address,
	// 	Phone: newUser.Phone,
	// 	CreatedAt: newUser.CreatedAt,
	// 	UpdatedAt: newUser.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS CREATE DATA",
		"data" : newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	var user entity.User

	result := db.First(&user, userId)

	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	Email: user.Email,
	// 	Name: user.Name,
	// 	Address: user.Address,
	// 	Phone: user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message" : "USER HAS FOUNDED",
		"data" : user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	// FIND USER
	var user entity.User

	userId := ctx.Params("userId")

	result := db.First(&user, userId)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
		})
	}

	// UPDATE DATA
	result = db.Debug().Model(&user).Updates(userRequest)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE DATA",
		})
	}

	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message" : "USER UPDATED",
		"data" : userResponse,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateEmailRequest)
	err := ctx.BodyParser(userRequest)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "BAD REQUEST",
		})
	}

	// FIND USER
	var user entity.User

	userId := ctx.Params("userId")

	result := db.First(&user, userId)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
		})
	}

	// UPDATE DATA
	if userRequest.Email != "" {
		err = validate.Struct(userRequest)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		user.Email = userRequest.Email
	}

	result = db.Debug().Save(&user)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "FAILED TO UPDATE EMAIL",
		})
	}

	userResponse := response.UpdateEmailUserResponse{
		Name: user.Name,
		Email: user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message" : "EMAIL UPDATED",
		"data" : userResponse,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	var user entity.User

	result := db.First(&user, userId)

	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message" : "USER NOT FOUND",
		})
	}

	// DELETE USER
	result = db.Debug().Delete(&user, userId)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "FAILED TO DELETE DATA",
		})
	}

	return ctx.JSON(fiber.Map{
		"message" : "SUCCESS DELETE USER",
	})
}