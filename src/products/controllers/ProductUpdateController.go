package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/products/services"
	"net/http"
)

func UpdateProduct(c *gin.Context) {
	var json dto.ProductUpdateDTO
	c.BindJSON(&json)

	service := services.ProductUpdateService{}
	service.UpdateProduct(c.Param("id"), &json)

	if service.ErrorValidation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"data": gin.H{
				"validation": service.ErrorValidation,
			},
		})
		return
	}

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
			"data": gin.H{
				"product": service.Product,
			},
		})
		return
	}

	c.JSON(http.StatusNotFound, nil)
	return
}
