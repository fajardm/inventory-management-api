package dto

type PurchaseCreateDTO struct {
	OrderQuantity  int     `form:"OrderQuantity" json:"OrderQuantity" validate:"required"`
	NumberReceived int     `form:"NumberReceived" json:"NumberReceived" validate:"gte=0,required"`
	Price          float32 `form:"Price" json:"Price" validate:"gte=0,required"`
	TotalPrice     float32 `form:"TotalPrice" json:"TotalPrice" validate:"gte=0,required"`
	Receipt        string  `form:"Receipt" json:"Receipt"`
	Note           string  `form:"Note" json:"Note"`
	ProductId      string  `form:"ProductId" json:"ProductId" validate:"required"`
}
