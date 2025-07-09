package main

import (
	taskModel "RestCrud/internal/task/model"
	userModel "RestCrud/internal/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=RestCrud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Auto-migrate tables
	if err := db.AutoMigrate(&userModel.User{}, &taskModel.Task{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("Database migration completed successfully")
}
