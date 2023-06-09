package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func openDatabase() (*gorm.DB, error) {
	// read database settings from environment
	databaseAddr := os.Getenv("WEBSERVICETEMPLATE_DB_URL")
	databaseUser := os.Getenv("WEBSERVICETEMPLATE_DB_USER")
	databasePassword := os.Getenv("WEBSERVICETEMPLATE_DB_PASSWORD")
	databaseDBName := os.Getenv("WEBSERVICETEMPLATE_DB_NAME")
	databasePort := os.Getenv("WEBSERVICETEMPLATE_DB_PORT")

	dbVariablesSet := (len(databaseAddr) != 0) && (len(databaseUser) != 0) && (len(databasePassword) != 0) && (len(databaseDBName) != 0) && (len(databasePort) != 0)

	if !dbVariablesSet {
		return nil, fmt.Errorf("db setting(s) missing")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", databaseAddr, databaseUser, databasePassword, databaseDBName, databasePort)
	dat, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return dat, nil
}
