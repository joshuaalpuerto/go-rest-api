package db

import (
	"database/sql"
	"fmt"

	"github.com/joshuaalpuerto/go-rest-api/config"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewDatabase(config *config.DBConf) (*Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgres{db: db}, nil
}

func (d *Postgres) Close() error {
	return d.db.Close()
}

func (d *Postgres) GetDB() *sql.DB {
	return d.db
}
