package db

import (
	"fmt"
	"strconv"

	"github.com/kurniawanxyz/crud-notes-go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Debug prints to check the configuration values
	fmt.Printf("DBUser: %s\n", config.ENV.DBUser)
	fmt.Printf("DBPass: %s\n", config.ENV.DBPass)
	fmt.Printf("DBHost: %s\n", config.ENV.DBHost)
	fmt.Printf("DBPort: %d\n", config.ENV.DBPort)
	fmt.Printf("DBName: %s\n", config.ENV.DBName)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ENV.DBUser, config.ENV.DBPass, config.ENV.DBHost, strconv.Itoa(config.ENV.DBPort), config.ENV.DBName)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}
