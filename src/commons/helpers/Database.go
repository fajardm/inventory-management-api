package helpers

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDatabase() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("sqlite3", "../../../../inventory.db")
	return db, err
}

func GetDatabase() *gorm.DB {
	return db
}
