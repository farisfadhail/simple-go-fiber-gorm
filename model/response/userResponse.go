package response

import "time"

type UserResponse struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"-" gorm:"column:password"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateEmailUserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	UpdatedAt time.Time `json:"updated_at"`
}