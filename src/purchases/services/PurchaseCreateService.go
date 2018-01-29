package services

import (
	"github.com/fajardm/inventories/src/purchases/domain"
	domain2 "github.com/fajardm/inventories/src/products/domain"
	"github.com/fajardm/inventories/src/purchases/dto"
	"github.com/fajardm/inventories/src/commons/helpers"
)

type PurchaseCreateService struct {
	helpers.Response
	EntityErrorMessage string
	Purchase           domain.Purchase
}

func (res *PurchaseCreateService) CreatePurchase(dto *dto.PurchaseCreateDTO) *PurchaseCreateService {
	err := helpers.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = helpers.GetValidationMessage(err)
	}

	db := helpers.GetDatabase()

	product := domain2.Product{}
	result := db.Where("id = ?", dto.Product).First(&product)

	if result.RecordNotFound() {
		res.EntityErrorMessage = "Product with id: " + dto.Product + " not found"
		return res
	} else if result.Error != nil {
		res.Error = result.Error.Error()
		return res
	}

	purchase := domain.Purchase{
		OrderQuantity:  dto.OrderQuantity,
		NumberReceived: dto.NumberReceived,
		Price:          dto.Price,
		TotalPrice:     dto.TotalPrice,
		Note:           dto.Note,
		Receipt:        dto.Receipt,
		Product:        product,
	}

	created := db.Create(&product)
	if created.Error == nil {
		res.Purchase = purchase
		return res
	} else {
		res.Error = created.Error.Error()
		return res
	}
}
