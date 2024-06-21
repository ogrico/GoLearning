package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn      = "geovany:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	Database = func() (db *gorm.DB) {
		if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			fmt.Println("Error durante la conexión a la db", err)
			panic(err)
		} else {
			fmt.Println("Conexión exitosa")
			return db
		}
	}
)
