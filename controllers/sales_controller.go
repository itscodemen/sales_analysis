package controllers

import (
	"net/http"
	"sales-analysis/services"
	"sales-analysis/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTotalRevenue(c *gin.Context) {
	startDate, err := time.Parse("2006-01-02", c.Query("start_date"))
	if err != nil {
		utils.Logger().Println("Invalid start date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", c.Query("end_date"))
	if err != nil {
		utils.Logger().Println("Invalid end date")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	totalRevenue, err := services.CalculateTotalRevenue(startDate, endDate)
	if err != nil {
		utils.Logger().Println("Failed to calculate total revenue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total revenue"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_revenue": totalRevenue})
}
