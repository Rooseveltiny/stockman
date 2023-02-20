package postgresutils

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"stockman/source/stockman_dbs/client/postgresql"
	logger "stockman/source/stockman_logger"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("failed to locate logfile")
		logger.L.Errorln(err)
	}
	c, _ := ioutil.ReadAll(f)
	fmt.Println(string(c))
}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePath []string) {}
