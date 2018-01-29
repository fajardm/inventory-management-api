package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/domain"
)

type ProductListService struct {
	helpers.Response
	Products []domain.Product
}

func (res *ProductListService) ListProduct() *ProductListService {
	db := helpers.GetDatabase()

	var products []domain.Product
	db.Find(&products)
	res.Products = products

	return res
}
