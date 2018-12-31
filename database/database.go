package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func CreatePool(connectionString string) (err error) {
	db, err = gorm.Open("mysql", connectionString)
	return
}

func GetConnection() *gorm.DB {
	return db
}

func Close() error {
	return db.Close()
}