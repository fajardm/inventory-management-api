package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fajardm/inventories/src/purchases/dto"
	"github.com/fajardm/inventories/src/purchases/services"
	"net/http"
)

func CreatePurchase(c *gin.Context) {
	var json dto.PurchaseCreateDTO
	c.BindJSON(&json)

	service := services.PurchaseCreateService{}
	service.CreatePurchase(&json)

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
			"purchase": service.Purchase,
		},
	})
	return
}
