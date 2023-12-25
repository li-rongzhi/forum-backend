package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

    "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

	"sample-go-app/internal/router"
	"sample-go-app/internal/models"
	"sample-go-app/internal/database"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	// Initialize the *gorm.DB object using gorm
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword)
	dbGORM, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database with GORM: %v", err)
	}

	// Auto migrate your schema using gorm
	if err = dbGORM.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate User with GORM: %v", err)
	}
	if err = dbGORM.AutoMigrate(&models.Category{}); err != nil {
		log.Fatalf("Failed to auto-migrate Category with GORM: %v", err)
	}
	if err = dbGORM.AutoMigrate(&models.Thread{}); err != nil {
		log.Fatalf("Failed to auto-migrate Thread with GORM: %v", err)
	}
	if err = dbGORM.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatalf("Failed to auto-migrate Comment with GORM: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to auto-migrate with GORM: %v", err)
	}

	database.DB = dbGORM

	r := router.Setup()
	// Setup CORS
	// corsHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"}, // Adjust the allowed origins to match your front-end app's URL
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: true,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// })
	// r.Use(corsHandler.Handler)

	fmt.Print("Listening on port 8000 at http://localhost:8000!")

	log.Fatalln(http.ListenAndServe(":8000", r))
}
