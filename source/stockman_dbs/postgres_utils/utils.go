package postgresutils

import (
	"context"
	"fmt"
	"io/ioutil"
	"stockman/source/stockman_dbs/client/postgresql"
	logger "stockman/source/stockman_logger"
	"strings"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to locate logfile")
		logger.L.Errorln(err)
	}
	SQLRAW := getRidOfScreening(string(content))
	p, err := client.Exec(ctx, SQLRAW)
	fmt.Println(p)
	fmt.Println(err)
}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePath []string) {}

func getRidOfScreening(i string) string {
	var screeningSymbols []string
	screeningSymbols = append(screeningSymbols, "\n")
	screeningSymbols = append(screeningSymbols, "\t")

	for _, s := range screeningSymbols {
		i = strings.ReplaceAll(i, s, "")
	}
	return i
}
