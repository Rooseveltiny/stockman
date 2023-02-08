package main

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestEvent(t *testing.T) {
	convey.Convey("Test event init", t, func() {
		event := NewEvent[string, string, string]()
		fmt.Println(event)
	})
}
