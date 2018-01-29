package services

import (
	"github.com/fajardm/inventories/src/sales/domain"
	domain2 "github.com/fajardm/inventories/src/products/domain"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/dto"
)

type SaleCreateService struct {
	helpers.Response
	EntityErrorMessage string
	Sale               domain.Sale
}

func (res *SaleCreateService) CreateSale(dto *dto.SaleCreateDTO) *SaleCreateService {
	db := helpers.GetDatabase()

	err := helpers.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = helpers.GetValidationMessage(err)
		return res
	}

	product := domain2.Product{}
	result := db.Where("id = ?", dto.ProductId).First(&product)

	if result.RecordNotFound() {
		res.EntityErrorMessage = "Product with id: " + dto.ProductId + " not found"
		return res
	} else if result.Error != nil {
		res.Error = result.Error.Error()
		return res
	}

	sale := domain.Sale{
		NumberShipped: dto.NumberShipped,
		Price:         dto.Price,
		TotalPrice:    dto.TotalPrice,
		Note:          dto.Note,
		ProductId:     product.ID,
	}

	created := db.Create(&sale)
	if created.Error == nil {
		product.Stock = product.Stock - sale.NumberShipped
		updatedProduct := db.Save(&product)

		if updatedProduct.Error != nil {
			res.Error = updatedProduct.Error.Error()
			return res
		}

		res.Sale = sale
		return res
	} else {
		res.Error = created.Error.Error()
		return res
	}
}
