package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/ByteForge-Systems/vpn-api/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL")
	return db, nil
}
