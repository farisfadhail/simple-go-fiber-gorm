package entity

import (
	"time"
)

type User struct {
	ID		 uint		`gorm:"primaryKey"`
	Name     string 	`json:"name"`
	Email    string 	`json:"email" gorm:"uniqueIndex"`
	Password string 	`json:"-" gorm:"column:password"`
	Address  string 	`json:"address"`
	Phone    string 	`json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// deleted_at akan tetap dimasukkan dengan nama column deleted_at namun tidak bisa dipanggil dengan deleted_at
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

//? gorm.Model equals here
// ID 		 uint 			`gorm:"primaryKey"`
// CreatedAt time.Time
// UpdatedAt time.Time
// DeletedAt gorm.DeletedAt `gorm:"index"`