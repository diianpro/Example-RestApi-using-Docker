package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type DbConfig struct {
	Port         int    `env:"APP_DB_PORT" envDefault:"8080"`
	DBName       string `env:"APP_DB_NAME" envDefault:"postgres"`
	Password     string `env:"APP_DB_PASSWORD"`
	Username     string `env:"APP_DB_USERNAME" envDefault:"postgres"`
	Host         string `envDefault:"localhost"`
	PgConnection *pgxpool.Pool
}

func (cfg *DbConfig) Init(ctx context.Context) error {
	connStr := fmt.Sprintf("host=localhost user=%s password=%s  dbname=%s port=%d sslmode=disable",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

	conn, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Store config: Can't connect to database: %s", err)
	}
	cfg.PgConnection = conn
	log.Info("connected to postgresql")

	return nil
}
