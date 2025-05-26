package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/LizaMeytner/foru/forum-service/config"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := sql.Open("postgres", cfg.Postgres.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Get migrations directory
	migrationsDir := "migrations"
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Fatalf("Migrations directory not found: %s", migrationsDir)
	}

	// Create migrations table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version bigint PRIMARY KEY,
			dirty boolean NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create migrations table: %v", err)
	}

	// Get current version
	var currentVersion int64
	err = db.QueryRow("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1").Scan(&currentVersion)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Failed to get current version: %v", err)
	}

	// Read migration files
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.up.sql"))
	if err != nil {
		log.Fatalf("Failed to read migration files: %v", err)
	}

	// Sort and apply migrations
	for _, file := range files {
		version := getVersionFromFilename(file)
		if version <= currentVersion {
			continue
		}

		// Read migration file
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", file, err)
		}

		// Start transaction
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Failed to start transaction: %v", err)
		}

		// Execute migration
		_, err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			log.Fatalf("Failed to execute migration %s: %v", file, err)
		}

		// Update version
		_, err = tx.Exec("INSERT INTO schema_migrations (version, dirty) VALUES ($1, false)", version)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Failed to update version: %v", err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			log.Fatalf("Failed to commit transaction: %v", err)
		}

		fmt.Printf("Applied migration %s\n", file)
	}

	fmt.Println("Migrations completed successfully")
}

func getVersionFromFilename(filename string) int64 {
	var version int64
	fmt.Sscanf(filepath.Base(filename), "%d_", &version)
	return version
}
