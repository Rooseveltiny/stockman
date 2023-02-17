package logger

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {

	convey.Convey("try to log smth", t, func() {
		L.Info("this is info message")
	})

}
