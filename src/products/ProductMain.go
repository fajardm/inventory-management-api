package products

import (
	"github.com/fajardm/inventories/src/commons/helpers"
	"github.com/fajardm/inventories/src/products/controllers"
)

func ProductRoutes() {
	router := helpers.GetRouter()

	router.GET("/products", controllers.ListProduct)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.ShowProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
