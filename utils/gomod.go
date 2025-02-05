package util

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	// db := config.Connection()
	// Database credentials
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Initialize generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./models", // Output directory for models
		ModelPkgPath: "models",
	})

	g.UseDB(db) // Use the connected database

	// Generate models from all tables
	g.GenerateAllTable()

	// Execute the generator
	g.Execute()
}
