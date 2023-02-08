package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEvent(t *testing.T) {
	Convey("test event init", t, func() {
		event := NewEvent[string, string, string]()
		So(event.ctx, ShouldNotBeNil)
		So(event.Data, ShouldEqual, "")
		So(event.Input, ShouldNotBeNil)
		So(event.Output, ShouldNotBeNil)
	})
}
