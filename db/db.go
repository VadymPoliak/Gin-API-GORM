package db

import (
	"project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func createTables(dbConn *gorm.DB) []error {

	dbConn.DropTableIfExists(&models.User{}).GetErrors()

	errs := dbConn.CreateTable(&models.User{}).GetErrors()

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ConnectDatabase() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:test123@192.168.200.178/testutf8?sslmode=disable")

	//dsn := "host=localhost user=postgres password=playwithyou dbname=test port=5432 sslmode=disable"
	//dbconn, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	// createTables(dbconn)
	dbconn.AutoMigrate(&models.User{}, &models.Children{}, &models.Campaign{}, &models.Event{}, &models.EventItem{})

	DB = dbconn

}
