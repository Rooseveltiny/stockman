package postgresql

import (
	"context"
	"fmt"
	"log"
	"restapi-lesson/internal/config"

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
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewClient(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	pool, err = pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Fatal("can't connect to db")
	}

	return pool, nil
}
