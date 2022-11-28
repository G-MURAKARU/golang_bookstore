package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// declaring database variable
var db *gorm.DB

func Connect() {
	// initialising the database
	database, err := gorm.Open("mysql", "gicheru:_murakaru.11@/bookstoreDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	// returning the DB for use
	return db
}
