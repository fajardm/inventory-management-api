package boot

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	productDomain "github.com/fajardm/inventories/src/products/domain"
	purchaseDomain "github.com/fajardm/inventories/src/purchases/domain"
	saleDomain "github.com/fajardm/inventories/src/sales/domain"
	reportDomain "github.com/fajardm/inventories/src/reports/domain"
)

func Migration() {
	db := helpers.GetDatabase()

	db.CreateTable(&productDomain.Product{})
	db.CreateTable(&purchaseDomain.Purchase{})
	db.CreateTable(&saleDomain.Sale{})

	reportDomain.CreateViewReportProductSale()
	reportDomain.CreateViewReportProductValue()
}
