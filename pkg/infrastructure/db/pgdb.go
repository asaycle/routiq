package db

import (
	"fmt"
	"log"

	"github.com/asaycle/routiq/pkg/lib/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL ドライバ
)

func NewPgDBFromCfg(cfg *config.DBConfig) (*sqlx.DB, error) {
	return NewPgDB(
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
}

func NewPgDB(host, port, user, password, dbname string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	log.Printf("NewPgdb: %s", dsn)
	return sqlx.Open("postgres", dsn)
}
