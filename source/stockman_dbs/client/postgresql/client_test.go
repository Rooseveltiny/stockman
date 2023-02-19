package postgresql

import (
	"context"
	logger "stockman/source/stockman_logger"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDBConfigInit(t *testing.T) {
	convey.Convey("test init db settings", t, func() {
		db_c := NewPostgresConfig()
		convey.So(db_c.Host, convey.ShouldEqual, "localhost")
	})
}

func TestPostgresClientInit(t *testing.T) {
	convey.Convey("test postgres client init process", t, func() {
		db_c := NewPostgresConfig()
		c, err := NewClient(context.TODO(), *db_c)
		if err != nil {
			logger.L.Errorln(err)
		}
		convey.So(c, convey.ShouldNotBeNil)
	})
}
