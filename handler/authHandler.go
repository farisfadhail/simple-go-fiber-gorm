package handler

import (
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	err := ctx.BodyParser(loginRequest)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	err = validate.Struct(loginRequest)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "BAD REQUEST",
			"error" : err.Error(),
		})
	}

	// CHECK AVAILABLE USER
	var user entity.User
	result := db.First(&user, "email = ?", loginRequest.Email)

	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	// CHECK VALIDATION PASSWORD
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	// GENERATE JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	//? Untuk menentukan role pada akun
	if user.Email == "udin@example.id" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	// claims["role"] = "user"
	
	token, err := utils.GenerateJwtToken(&claims)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "WRONG CREDENTIAL",
		})
	}

	return ctx.JSON(fiber.Map{
		"token" : token,
	})
}