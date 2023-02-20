package postgresutils

import (
	"context"
	"fmt"
	"os"
	"stockman/source/stockman_dbs/client/postgresql"
	logger "stockman/source/stockman_logger"
	"strings"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) (ok bool) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to locate logfile")
		logger.L.Errorln(err)
	}
	SQLRAW := getRidOfScreening(string(content))
	_, err_sql := client.Exec(ctx, SQLRAW)
	if err_sql != nil {
		logger.L.Errorln(err_sql)
		return false
	}
	return true
}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePath []string) {

}

func getRidOfScreening(i string) string {
	var screeningSymbols []string
	screeningSymbols = append(screeningSymbols, "\n")
	screeningSymbols = append(screeningSymbols, "\t")

	for _, s := range screeningSymbols {
		i = strings.ReplaceAll(i, s, "")
	}
	return i
}
