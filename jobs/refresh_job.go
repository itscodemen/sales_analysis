package jobs

import (
	"log"
	"sales-analysis/scripts"
	"time"

	"gorm.io/gorm"
)

func StartDataRefreshJob(filePath string, interval time.Duration, db *gorm.DB) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				refreshData(filePath, db)
			}
		}
	}()
}

func refreshData(filePath string, db *gorm.DB) {
	log.Println("Starting data refresh...")

	// Load data from CSV and insert/update into the database
	scripts.LoadCSVData(filePath, db)

	log.Println("Data refresh completed at", time.Now())
}
