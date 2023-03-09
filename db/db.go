package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func Init() {
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
}
