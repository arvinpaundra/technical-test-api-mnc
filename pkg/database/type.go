package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the database conn configuration
type (
	dbConfig struct {
		dsn string
	}
)

// Connect to postgres with the input configuration
func (conf dbConfig) connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conf.dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}
