package request

type ImageCreateRequest struct {
	CategoryId uint `json:"category_id" form:"category_id" validate:"required"`
}