package scripts

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sales-analysis/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func LoadCSVData(filePath string, db *gorm.DB) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	for _, record := range records[1:] {
		orderID := record[0]
		productID := record[1]
		customerID := record[2]
		productName := record[3]
		category := record[4]
		dateOfSale, _ := time.Parse("2006-01-02", record[5])
		quantitySold, _ := strconv.Atoi(record[6])
		unitPrice, _ := strconv.ParseFloat(record[7], 64)
		discount, _ := strconv.ParseFloat(record[8], 64)
		shippingCost, _ := strconv.ParseFloat(record[9], 64)
		paymentMethod := record[10]
		customerName := record[11]
		customerEmail := record[12]
		customerAddress := record[13]

		product := models.Product{
			ProductID:   productID,
			ProductName: productName,
			Category:    category,
		}

		customer := models.Customer{
			CustomerID:      customerID,
			CustomerName:    customerName,
			CustomerEmail:   customerEmail,
			CustomerAddress: customerAddress,
		}

		order := models.Order{
			OrderID:       orderID,
			ProductID:     productID,
			CustomerID:    customerID,
			DateOfSale:    dateOfSale,
			QuantitySold:  quantitySold,
			UnitPrice:     unitPrice,
			Discount:      discount,
			ShippingCost:  shippingCost,
			PaymentMethod: paymentMethod,
		}

		// Use FirstOrCreate for product and customer to avoid duplicates
		if err := db.Where(&models.Product{ProductID: productID}).FirstOrCreate(&product).Error; err != nil {
			log.Fatalf("Failed to upsert product: %v", err)
		}

		if err := db.Where(&models.Customer{CustomerID: customerID}).FirstOrCreate(&customer).Error; err != nil {
			log.Fatalf("Failed to upsert customer: %v", err)
		}

		// Use Create to insert order, or Update if it already exists
		if err := db.Where(&models.Order{OrderID: orderID}).Assign(order).FirstOrCreate(&order).Error; err != nil {
			log.Fatalf("Failed to upsert order: %v", err)
		}
	}
	fmt.Println("CSV data loaded successfully")
}
