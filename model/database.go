package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db instance that will be used for db operations
var db *gorm.DB

func ConnectToDB() *gorm.DB {
	configurations := "tong:1209@/buzzer_beater?charset=utf8mb4&parseTime=True&loc=Local"
	connectedDB, err := gorm.Open("mysql", configurations)

	if err != nil {
		fmt.Println("database connection failed: " + err.Error())
	} else {
		log.Println("database connection succedssed")
	}

	// set table name to be singular, for instnce: User instead of Users
	connectedDB.SingularTable(true)
	return connectedDB
}

// set db variable to current connection so that it can used else where
func SetDB(database *gorm.DB) {
	db = database
}
