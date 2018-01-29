package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/products/services"
	"net/http"
)

func ListProduct(c *gin.Context) {
	service := services.ProductListService{}
	service.ListProduct()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"products": service.Products,
		},
	})
	return
}
