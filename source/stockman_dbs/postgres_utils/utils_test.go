package postgresutils

import (
	"context"
	"fmt"
	"stockman/source"
	"stockman/source/stockman_dbs/client/postgresql"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestRunSQLFunc(t *testing.T) {
	convey.Convey("test sql raw exec", t, func() {
		cfg := *postgresql.NewPostgresConfig()
		c, _ := postgresql.NewClient(context.TODO(), cfg)
		ok := RunSQLFile(context.TODO(), c, "test_files/sql_1.sql")
		convey.So(ok, convey.ShouldBeTrue)
		convey.Convey("test sql raw write exec", func() {
			SQLRAW := `
			INSERT INTO test_table (username)
			VALUES ($1)`
			_, err := c.Exec(context.TODO(), SQLRAW, "Saveliy")
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestApplyAllPreparedSQL(t *testing.T) {
	convey.Convey("init all file path", t, func() {
		cfg := postgresql.NewPostgresConfig(source.DB_TEST_YAML)
		l := getListOfPostgresSQLFiles(cfg.SqlFolder)
		convey.So(strings.Contains(l[len(l)-1], ".sql"), convey.ShouldBeTrue)
		convey.Convey("apply all postgres tables", func() {
			ctx := context.TODO()
			postgresClient, _ := postgresql.NewClient(ctx, *postgresql.NewPostgresConfig(source.DB_TEST_YAML))
			err := RunPostgresSQL(ctx, postgresClient)
			fmt.Println(err)
		})
	})
}
