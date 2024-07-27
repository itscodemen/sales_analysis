package main

import (
	"sales-analysis/config"
	"sales-analysis/jobs"
	"sales-analysis/models"
	"sales-analysis/routes"
	"sales-analysis/scripts"
	"sales-analysis/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	models.ConnectDatabase(cfg)

	// Start the data refresh job
	filePath := "sample.csv"
	scripts.LoadCSVData(filePath, models.DB)
	refreshInterval := 24 * time.Hour // Refresh every 24 hours
	jobs.StartDataRefreshJob(filePath, refreshInterval, models.DB)

	r := gin.Default()
	routes.RegisterRoutes(r)

	utils.Logger().Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		utils.Logger().Fatal("Failed to start server: ", err)
	}
}
