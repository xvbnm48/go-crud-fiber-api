package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-rest-fiber)"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})

	DB = db
}
