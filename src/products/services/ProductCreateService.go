package services

import (
	"github.com/fajardm/inventories/src/products/domain"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/commons/helpers"
)

type ProductCreateService struct {
	helpers.Response
	Product domain.Product
}

func (res *ProductCreateService) CreateProduct(dto *dto.ProductCreateDTO) *ProductCreateService {
	err := helpers.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = helpers.GetValidationMessage(err)
		return res
	}

	db := helpers.GetDatabase()
	product := domain.Product{SKU: dto.SKU, Name: dto.Name, Stock: dto.Stock}

	created := db.Create(&product)
	if created.Error == nil {
		res.Product = product
		return res
	} else {
		res.Error = created.Error.Error()
		return res
	}
}
