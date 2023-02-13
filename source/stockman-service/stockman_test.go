package main

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestyDTO struct {
	TestName string `json:"test_name"`
}

func TestEvent(t *testing.T) {
	Convey("test event init", t, func() {
		f := func(ctx context.Context, e *Event) {
			fmt.Println("test print ln")
			e.NotifyOutputChanged()
		}
		e := NewEvent(f)
		So(e.event, ShouldNotBeNil)
		So(e.ctx, ShouldNotBeNil)
		So(e.mu, ShouldNotBeNil)
		So(e.input, ShouldBeEmpty)
		So(e.output, ShouldBeEmpty)
		So(e.outputChanged, ShouldNotBeNil)
		Convey("test event run", func() {
			e.RunEvent()
			<-e.OnOutputChanged()
			Convey("test sending input data", func() {
				testName := "Robert"
				e.SetOutput(TestyDTO{TestName: testName})
				output := &TestyDTO{}
				e.LoadOutput(output)
				So(output.TestName, ShouldEqual, testName)
			})
		})
	})
}
