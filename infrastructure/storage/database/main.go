package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	dialect := "mysql"
	username := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "zews-buzz"

	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(dialect, conn)

	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)
	return db
}

func Disconnect(db *gorm.DB) {
	defer db.Close()
}
