package postgresutils

import "testing"

func TestRunSQLFunc(t *testing.T) {
	RunSQLFile("test_files/sql_1.sql")
}
