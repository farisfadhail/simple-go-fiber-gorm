package routes

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath + "/public/asset")
	
	api := app.Group("/api")

	// Auth Routes
	api.Post("/login", handler.LoginHandler).Name("login")
	
	// User Routes
	user := api.Group("/user", middleware.Authenticated, middleware.IsAdmin)
	user.Get("/", handler.UserHandlerGetAll).Name("user.index")
	user.Post("/", handler.UserHandlerCreate).Name("user.store")
	user.Get("/:userId", handler.UserHandlerGetById).Name("user.show")
	user.Put("/:userId", handler.UserHandlerUpdate).Name("user.update")
	user.Put("/:userId/update-email", handler.UserHandlerUpdateEmail).Name("user.emailUpdate")
	user.Delete("/:userId", handler.UserHandlerDelete).Name("user.destroy")

	// Book Routes
	book := api.Group("/book")
	book.Get("/", handler.BookHandlerGetAll).Name("book.index")
	book.Post("/", utils.HandleSingleFile, handler.BookHandlerCreate).Name("book.store")

	// Routes
	api.Post("/gallery", utils.HandleMultipleFile, handler.ImageGalleriesHandlerCreate).Name("gallery")
	api.Delete("/gallery/:imageId", handler.ImageGalleriesHandlerDestroy).Name("gallery.destroy")
}