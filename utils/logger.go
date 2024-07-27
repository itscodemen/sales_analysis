package utils

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "sales-analysis: ", log.LstdFlags)
}
