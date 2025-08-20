package test

import (
	"gogym/tracker/internal/db"
	"gogym/tracker/internal/shared"
	"log"
	"testing"
)

func TestDB(m *testing.T) {
	// Initialize the config
	config := shared.NewConfig()
	config.Load("../.env")
	log.Printf("DB User: %s\n", config.DBUser)
	log.Printf("DB Password: %s\n", config.DBPassword)
	log.Printf("DB Name: %s\n", config.DBName)
	log.Printf("DB Host: %s\n", config.DBHost)
	log.Printf("DB Port: %s\n", config.DBPort)

	dbConn := db.NewDBConnection()
	log.Printf("DB Connection: %+v\n", dbConn)

	log.Println("Running a basic test...")
}
