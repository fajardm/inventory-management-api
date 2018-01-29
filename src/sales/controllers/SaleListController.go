package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fajardm/inventories/src/sales/services"
)

func ListSales(c *gin.Context) {
	service := services.SaleListService{}
	service.ListSales()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"sales": service.Sales,
		},
	})
	return
}
