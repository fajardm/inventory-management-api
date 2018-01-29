package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/products/dto"
	"github.com/fajardm/inventories/src/products/services"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var json dto.ProductCreateDTO
	c.BindJSON(&json)

	service := services.ProductCreateService{}
	service.CreateProduct(&json)

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

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"product": service.Product,
		},
	})
	return
}
