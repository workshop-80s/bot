package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect(config map[string]string) *gorm.DB {
	dialect := config["dialect"]
	username := config["username"]
	password := config["password"]
	host := config["host"]
	port := config["port"]
	database := config["database"]

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
