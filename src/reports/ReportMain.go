package reports

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/reports/controllers"
)

func ReportRoutes() {
	router := helpers.GetRouter()

	router.GET("/reports/product_sales", controllers.ReportProductSale)
	router.GET("/reports/product_values", controllers.ReportProductValue)
}
