package postgresutils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"stockman/source"
	"stockman/source/stockman_dbs/client/postgresql"
	logger "stockman/source/stockman_logger"
	"strings"
)

func RunSQLFile(ctx context.Context, client postgresql.Client, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to locate logfile")
		logger.L.Errorln(err)
	}
	SQLRAW := getRidOfScreening(string(content))
	_, err_sql := client.Exec(ctx, SQLRAW)
	if err_sql != nil {
		logger.L.Errorln(err_sql)
		return err_sql
	}
	return nil
}

func RunSQLFiles(ctx context.Context, client postgresql.Client, filePaths []string) error {
	for _, f := range filePaths {
		err := RunSQLFile(ctx, client, f)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
Function applies all sql tables and commands earlier prepared
*/
func RunPostgresSQL(ctx context.Context, client postgresql.Client, dbConfigPath string) error {
	allSqlFilesToPerform := getListOfPostgresSQLFiles(postgresql.NewPostgresConfig(dbConfigPath).SqlFolder)
	err := RunSQLFiles(ctx, client, allSqlFilesToPerform)
	if err != nil {
		return err
	}
	return nil
}

func PrepareTestPostgresSQL(ctx context.Context, client postgresql.Client) error {
	return RunPostgresSQL(ctx, client, source.DB_TEST_YAML)
}

func DropPreparedTestPostgresSQL(ctx context.Context, client postgresql.Client) error {
	q :=
		`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`
	_, err := client.Exec(ctx, q)
	return err
}

func getListOfPostgresSQLFiles(baseDir string) []string {

	sqlFilePaths := make([]string, 0)
	subDirs := make([]string, 0)

	/* retrieve base directory */
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		logger.L.Errorln(err)
	}

	/* collecting all directories with files */
	for _, e := range entries {
		if e.IsDir() {
			subDirs = append(subDirs, filepath.Join(baseDir, e.Name()))
		}
	}

	/* collecting all sql files */
	for _, subDir := range subDirs {
		sqlFiles, err := os.ReadDir(subDir)
		if err != nil {
			logger.L.Errorln(err)
		}
		for _, sqlFile := range sqlFiles {
			if !sqlFile.IsDir() {
				if strings.Contains(sqlFile.Name(), ".sql") {
					sqlFilePaths = append(sqlFilePaths, filepath.Join(subDir, sqlFile.Name()))
				}
			}
		}
	}

	return sqlFilePaths
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
