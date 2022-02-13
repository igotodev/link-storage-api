package psql

import (
	"database/sql"
	"fmt"
	"link-storage-api/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

func NewPSLQ(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresDB.Host, cfg.PostgresDB.Port, cfg.PostgresDB.User, cfg.PostgresDB.Password, cfg.PostgresDB.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
