package sales

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/sales/controllers"
)

func SaleRoutes() {
	router := helpers.GetRouter()

	router.GET("/sales", controllers.ListSales)
	router.POST("/sales", controllers.CreateSale)
}
