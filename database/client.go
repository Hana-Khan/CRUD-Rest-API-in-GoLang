package database

import (
	"crud-rest-api-in-golang/entities"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//  Defines an instance of the database 
// and an error variable that will be used within other functions.
var Instance *gorm.DB
var err error

// This function basically attempts to connect to the database via GORM helpers using the connection string provided in our config.json. 
// Once connected, the variable Instance will be able to access the database to perform operations.
func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}
// It’s kinda important to make sure that the Entities in concern exist as tables in the connected database. 
// This particular method ensures that a table named products is created on the connected database.
func Migrate(){	
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database Migration Completed...")
}