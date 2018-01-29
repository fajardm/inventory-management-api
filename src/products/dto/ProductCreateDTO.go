package dto

type ProductCreateDTO struct {
	SKU   string `form:"SKU" json:"SKU" validate:"required"`
	Name  string `form:"Name" json:"Name" validate:"required"`
	Stock int    `form:"Stock" json:"Stock" validate:"gte=0,required"`
}
