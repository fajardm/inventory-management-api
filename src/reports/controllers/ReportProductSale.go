package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fajardm/inventories/src/reports/services"
	"time"
)

func ReportProductSale(c *gin.Context) {
	service := services.ReportProductSaleService{}
	service.ReportProductSale()

	if service.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": service.Error,
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=report_product_sale_"+time.Now().String()+".csv")
	c.Data(http.StatusOK, "text/csv", service.Buffer.Bytes())
	return
}
