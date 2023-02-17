package postgresql

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDBConfigInit(t *testing.T) {
	convey.Convey("test init db settings", t, func() {
		db_c := NewPostgresConfig()
		convey.So(db_c.Host, convey.ShouldEqual, "localhost")
	})
}
