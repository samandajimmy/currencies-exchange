package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func InitialMigration() {
	db, err = gorm.Open("postgres", getConnParams())

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&FromTo{}, &Rate{})
}

func getConnParams() string {
	connection := "host=" + os.Getenv("DB_HOST") + " dbname=" + os.Getenv("DB_NAME")

	if os.Getenv("DB_USER") != "" {
		connection += " user=" + os.Getenv("DB_USER")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		connection += " password=" + os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_PORT") != "" {
		connection += " port=" + os.Getenv("DB_PORT")
	}

	if os.Getenv("DB_SSLMODE") != "" {
		connection += " sslmode=" + os.Getenv("DB_SSLMODE")
	}

	return connection
}
