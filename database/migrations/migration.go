package migrations

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"log"
)

func RunMigration() {
	db := database.DatabaseInit()

	// db.Migrator().DropTable(&entity.User{})
	db.Migrator().DropTable(&entity.Book{})
	// db.Migrator().DropTable(&entity.Category{})
	// db.Migrator().DropTable(&entity.ImageGallery{})
	
	err := db.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.ImageGallery{})
	
	if err != nil {
		log.Println(err)
	}
	
	// fmt.Println("Database Migrated Fresh")
	fmt.Println("Database Migrated")
}