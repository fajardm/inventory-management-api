package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products"
	"github.com/fajardm/inventories/src/purchases"
	"github.com/fajardm/inventories/src/sales"
	"github.com/fajardm/inventories/src/reports"
)

func SetupMain() *gin.Engine {
	helpers.SetDatabase()
	Migration()
	Seeder()
	helpers.SetValidator()

	router := helpers.SetRouter()

	products.ProductRoutes()
	purchases.PurchaseRoutes()
	sales.SaleRoutes()
	reports.ReportRoutes()

	return router
}
