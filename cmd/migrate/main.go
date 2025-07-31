package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joshuaalpuerto/go-rest-api/config"
	"github.com/joshuaalpuerto/go-rest-api/internal/infra/db"
)

func main() {
	cfg := config.New()
	database, err := db.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Run migrations
	if err := runMigrations(database.GetDB()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrations completed successfully!")
}

func runMigrations(db *sql.DB) error {
	migrationsDir := "internals/infra/migrations/postgres"
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute %s: %w", file.Name(), err)
		}

		fmt.Printf("Executed migration: %s\n", file.Name())
	}

	return nil
}
