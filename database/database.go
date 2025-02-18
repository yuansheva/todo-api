package database

import (
	"fmt"
	"log"
	"todo-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "postgresql://postgres.cnqdokspjafgydavxzoi:@Sheva210@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")

	// Auto migrate models
	DB.AutoMigrate(&models.Todo{})
}
