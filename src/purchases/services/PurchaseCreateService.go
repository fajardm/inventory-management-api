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
	result := db.Where("id = ?", dto.ProductId).First(&product)

	if result.RecordNotFound() {
		res.EntityErrorMessage = "Product with id: " + dto.ProductId + " not found"
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
		ProductId:      product.ID,
	}

	created := db.Create(&purchase)
	if created.Error == nil {
		product.Stock = purchase.OrderQuantity + product.Stock
		updatedProduct := db.Save(&product)

		if updatedProduct.Error != nil {
			res.Error = updatedProduct.Error.Error()
			return res
		}

		res.Purchase = purchase
		return res
	} else {
		res.Error = created.Error.Error()
		return res
	}
}
