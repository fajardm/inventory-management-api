package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/domain"
)

type SaleListService struct {
	helpers.Response
	Sales []domain.Sale
}

func (res *SaleListService) ListSales() *SaleListService {
	db := helpers.GetDatabase()

	db.Find(&res.Sales)

	return res
}
