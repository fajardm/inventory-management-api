package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"github.com/fajardm/inventories/src/reports/services"
)

func ReportProductValue(c *gin.Context) {
	service := services.ReportProductValueService{}
	service.ReportProductValue()

	if service.Error != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": service.Error,
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=report_product_value_"+time.Now().String()+".csv")
	c.Data(http.StatusOK, "text/csv", service.Buffer.Bytes())
	return
}
