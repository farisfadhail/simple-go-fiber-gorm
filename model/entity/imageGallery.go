package entity

import "time"

type ImageGallery struct {
	ID        	uint      	`gorm:"primaryKey"`
	CategoryId 	uint		`json:"category_id"`
	Image     	string		`json:"image"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}