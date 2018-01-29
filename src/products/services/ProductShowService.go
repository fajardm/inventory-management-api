package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/domain"
)

type ProductShowService struct {
	helpers.Response
	Product domain.Product
}

func (res *ProductShowService) ShowProduct(id string) *ProductShowService {
	db := helpers.GetDatabase()

	product := domain.Product{}
	result := db.Where("id = ?", id).First(&product)

	if result.RecordNotFound() {
		return res
	} else if result.Error != nil {
		res.Error = result.Error.Error()
		return res
	} else {
		res.IsFound = true
		res.Product = product
		return res
	}
}
