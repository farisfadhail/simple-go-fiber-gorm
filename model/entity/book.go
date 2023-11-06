package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"uniqueIndex"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}