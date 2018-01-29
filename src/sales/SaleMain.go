package sales

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/controllers"
)

func SaleRoutes() {
	router := helpers.GetRouter()

	router.POST("/sales", controllers.CreateSale)
}