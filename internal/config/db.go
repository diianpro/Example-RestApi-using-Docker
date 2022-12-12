package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type DbConfig struct {
	Port         int    `env:"APP_DB_PORT" envDefault:"5433"`
	DBName       string `env:"APP_DB_NAME" envDefault:"postgres"`
	Password     string `env:"APP_DB_PASSWORD"`
	Username     string `env:"APP_DB_USERNAME" envDefault:"postgres"`
	Host         string `envDefault:"localhost"`
	PgConnection *pgxpool.Pool
}

func (cfg *DbConfig) Init(ctx context.Context) error {
	connStr := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Username, cfg.Password, cfg.DBName, cfg.Port)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(config.ConnConfig.ConnString())
	conn, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("Store config: Can't connect to database: %s", err)
	}
	cfg.PgConnection = conn
	log.Info("connected to postgresql")

	return nil
}
