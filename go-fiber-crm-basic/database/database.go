// package database

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// var (
// 	DBConn *gorm.DB
// )

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func ConnectDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}
