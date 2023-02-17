package postgresql

import (
	"context"
	"fmt"
	logger "stockman/source/stockman_logger"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

/*
Config struct to init postgres client connection to db
*/
type PostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func (sc *PostgresConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
}

func NewPostgresConfig() *PostgresConfig {
	cfg := PostgresConfig{}
	err := cleanenv.ReadConfig("db_settings.yaml", &cfg)
	if err != nil {
		logger.L.Errorln(err)
	}
	return &cfg
}

func NewClient(ctx context.Context, sc PostgresConfig) (pool *pgxpool.Pool, err error) {
	dsn := sc.DSN()

	pool, err = pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err != nil {
		logger.L.Errorln("can't connect to db")
	}

	return pool, nil
}
