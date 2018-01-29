package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/domain"
	"github.com/fajardm/inventories/src/products/dto"
)

type ProductUpdateService struct {
	helpers.Response
	Product domain.Product
}

func (res *ProductUpdateService) UpdateProduct(id string, dto *dto.ProductUpdateDTO) *ProductUpdateService {
	db := helpers.GetDatabase()

	err := helpers.GetValidator().Struct(dto)
	if err != nil {
		res.ErrorValidation = helpers.GetValidationMessage(err)
		return res
	}

	product := domain.Product{}
	result := db.Where("id = ?", id).First(&product)

	if result.RecordNotFound() {
		return res
	} else if result.Error != nil {
		res.Error = result.Error.Error()
		return res
	}

	product.SKU = dto.SKU
	product.Name = dto.Name
	product.Stock = dto.Stock
	updated := db.Save(&product)

	if updated.Error == nil {
		res.IsFound = true
		res.Product = product
		return res
	} else {
		res.Error = updated.Error.Error()
		return res
	}
}
