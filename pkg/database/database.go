package database

import (
	"log"
	"sync"

	"github.com/arvinpaundra/technical-test-api-mnc/config"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	dbErr  error
	once   sync.Once
)

func createConnection() {
	// Create database configuration information
	dbCfg := dbConfig{
		dsn: config.GetPostgresDSN(),
	}

	// Create only one database Connection, not the same as database TCP connection
	once.Do(func() {
		dbConn, dbErr = dbCfg.connect()
		if dbErr != nil {
			log.Fatalf("failed connected to database: %s", dbErr.Error())
		}
	})

	log.Println("connected to database")
}

func GetConnection() *gorm.DB {
	// Check db connection, if exist return the memory address of the db connection
	if dbConn == nil {
		createConnection()
	}
	return dbConn
}
