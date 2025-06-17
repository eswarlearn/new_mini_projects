package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectDB establishes a connection to the database
var (
	db *gorm.DB
)

func Connect() { //help to connect to the database
	d, err := gorm.Open("mysql", "go_user:learningGo$99@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local") //("db","username:password@tablename?)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB { // GetDB returns the database connection
	return db
}
