package postgresutils

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
	"testing"
)

func TestRunSQLFunc(t *testing.T) {
	cfg := *postgresql.NewPostgresConfig()
	c, _ := postgresql.NewClient(context.TODO(), cfg)
	RunSQLFile(context.TODO(), c, "test_files/sql_1.sql")
}
