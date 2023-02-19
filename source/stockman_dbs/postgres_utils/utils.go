package postgresutils

import (
	"context"
	"fmt"
	"os"
	"stockman/source/stockman_dbs/client/postgresql"
	logger "stockman/source/stockman_logger"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("failed to locate logfile")
		logger.L.Errorln(err)
	}

	fmt.Println(f)

}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePath []string) {}
