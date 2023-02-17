package logger

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {

	convey.Convey("try to log smth", t, func() {
		f, err := os.OpenFile("test_logging.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Failed to create logfile" + "system_logging.log")
			panic(err)
		}
		defer f.Close()
		L.Out = io.MultiWriter(f, os.Stdout)
		L.Info("this is info message")
	})

}
