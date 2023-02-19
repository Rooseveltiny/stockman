package postgresutils

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) {}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePath []string) {}
