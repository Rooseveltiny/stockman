package postgresutils

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
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
