package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/R2Remote/ChronoGo/internal/domain/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Ensure data directory exists
	dbFile := "data/chrono.db"
	if err := os.MkdirAll(filepath.Dir(dbFile), 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// 2. Connect to SQLite
	log.Printf("Connecting to SQLite database: %s...", dbFile)
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}
	log.Println("Connected to SQLite.")

	// 3. AutoMigrate Schema
	log.Println("Migrating schema...")
	if err := db.AutoMigrate(&entity.Job{}); err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}

	log.Println("Migration complete! Tables 'jobs' are ready.")
}
