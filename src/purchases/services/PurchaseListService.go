package services

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/purchases/domain"
)

type PurchaseListService struct {
	helpers.Response
	Purchases []domain.Purchase
}

func (res *PurchaseListService) ListPurchases() *PurchaseListService {
	db := helpers.GetDatabase()

	db.Find(&res.Purchases)

	return res
}
