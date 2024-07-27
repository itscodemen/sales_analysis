package services

import (
	"sales-analysis/models"
	"time"
)

func CalculateTotalRevenue(startDate, endDate time.Time) (float64, error) {
	var totalRevenue float64
	err := models.DB.Model(&models.Order{}).
		Where("date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Select("SUM(quantity_sold * unit_price * (1 - discount))").
		Scan(&totalRevenue).Error

	if err != nil {
		return 0, err
	}
	return totalRevenue, nil
}
