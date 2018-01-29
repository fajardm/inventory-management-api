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
	result := db.Where("id = ?", dto.Product).First(&product)

	if result.RecordNotFound() {
		res.EntityErrorMessage = "Product with id: " + dto.Product + " not found"
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
		Product:       product,
	}

	created := db.Create(&product)
	if created.Error == nil {
		res.Sale = sale
		return res
	} else {
		res.Error = created.Error.Error()
		return res
	}
}
