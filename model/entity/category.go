package entity

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	ImageGalleries []ImageGallery `json:"image_galleries"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}