package database

import (
	"fmt"
	"os"
	"time"

	"github.com/ZiplEix/super_snake/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() error {
	fmt.Println("Connecting to database")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
	// 	os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB"),
	// 	os.Getenv("POSTGRES_PORT"),
	// 	"Europe/Paris",
	// )

	dsn := os.Getenv("POSTGRES_URL")

	var err error
	for i := 0; i < 5; i++ {
		Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Connexion réussie
			fmt.Println("Successfully connected to database")
			return nil
		}

		fmt.Printf("Failed to connect to database, retrying... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}

	// Si toutes les tentatives échouent
	return fmt.Errorf("failed to connect to database after 5 attempts: %w", err)
}

func Migrate() error {
	fmt.Println("Migrating database...")

	err := Db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
