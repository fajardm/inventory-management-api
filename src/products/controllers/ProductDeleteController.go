package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/products/services"
	"net/http"
)

func DeleteProduct(c *gin.Context) {
	service := services.ProductDeleteService{}
	service.DeleteProduct(c.Param("id"))

	if service.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": service.Error,
		})
		return
	}

	if service.IsFound {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusNotFound, nil)
	return
}
