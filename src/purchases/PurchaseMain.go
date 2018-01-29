package purchases

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/purchases/controllers"
)

func PurchaseRoutes() {
	router := helpers.GetRouter()

	//router.GET("/products", controllers.ListProduct)
	router.POST("/purchases", controllers.CreatePurchase)
}
