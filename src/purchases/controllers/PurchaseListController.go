package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fajardm/inventories/src/purchases/services"
)

func ListPurchases(c *gin.Context) {
	service := services.PurchaseListService{}
	service.ListPurchases()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"purchases": service.Purchases,
		},
	})
	return
}
