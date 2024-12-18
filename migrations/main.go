package main

import (
	"fmt"
	"hr-system/common/config"
	"hr-system/common/dao/models"
	"hr-system/seeddata"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("migration start")

	// 資料庫連接
	dsn := config.Get().Mysql.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Employee{}, &models.Department{}, &models.Position{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migration completed successfully!")
	fmt.Println("migration end")

	seeddata.Gen(db)
}
