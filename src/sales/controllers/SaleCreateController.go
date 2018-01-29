package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/sales/dto"
	"github.com/fajardm/inventories/src/sales/services"
	"net/http"
)

func CreateSale(c *gin.Context) {
	var json dto.SaleCreateDTO
	c.BindJSON(&json)

	service := services.SaleCreateService{}
	service.CreateSale(&json)

	if service.ErrorValidation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"data": gin.H{
				"validation": service.ErrorValidation,
			},
		})
		return
	}

	if service.EntityErrorMessage != "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": service.EntityErrorMessage,
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
			"sale": service.Sale,
		},
	})
	return
}
