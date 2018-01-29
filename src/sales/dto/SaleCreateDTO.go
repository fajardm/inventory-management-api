package dto

type SaleCreateDTO struct {
	NumberShipped int     `form:"NumberShipped" json:"NumberShipped" validate:"gte=0,required"`
	Price         float32 `form:"Price" json:"Price" validate:"gte=0,required"`
	TotalPrice    float32 `form:"TotalPrice" json:"TotalPrice" validate:"gte=0,required"`
	Note          string  `form:"Note" json:"Note"`
	ProductId     string  `form:"ProductId" json:"ProductId" validate:"required"`
}
