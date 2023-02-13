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

type OutputDTO struct {
	TestFieldOfOutput string `json:"test_field_of_output"`
}

var outputFn = func(ctx context.Context, e *Event) {
	e.SetOutput(OutputDTO{TestFieldOfOutput: "Some data"})
	e.NotifyOutputChanged()
}

var f EventFn = func(ctx context.Context, e *Event) {
	fmt.Println("test print ln")
	e.NotifyOutputChanged()
}

func TestEvent(t *testing.T) {
	Convey("test event init", t, func() {
		e := NewEvent(f)
		So(e.event, ShouldNotBeNil)
		So(e.ctx, ShouldNotBeNil)
		// So(e.mu, ShouldNotBeNil)
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

func TestEventsList(t *testing.T) {
	Convey("test events list init", t, func() {
		el := NewEventsList()
		So(el.Size(), ShouldBeZeroValue)
		Convey("test append method", func() {
			e := NewEvent(f)
			e1 := NewEvent(f)
			e2 := NewEvent(f)
			e3 := NewEvent(f)
			e4 := NewEvent(f)
			el.AppendEvent(e)
			el.AppendEvent(e1)
			el.AppendEvent(e2)
			el.AppendEvent(e3)
			el.AppendEvent(e4)
			So(el.Size(), ShouldEqual, 5)
			Convey("test get last method of list", func() {
				lastUUID := el.Events[len(el.Events)-1].uuid
				getLastEvent := el.PopEvent()
				So(lastUUID, ShouldEqual, getLastEvent.uuid)
			})
		})
	})
}

func TestEventsPool(t *testing.T) {
	Convey("test events pool", t, func() {
		e := NewEvent(f)
		ep := NewEventsPool()
		ep.AppendEvent(e)
		ev := ep.PopEvent()
		So(e, ShouldEqual, ev)
	})
}

func TestEventsManager(t *testing.T) {
	Convey("test events manager init", t, func() {
		em := NewSystemEventsManager()
		So(em.Events.Size(), ShouldBeZeroValue)
		So(em.BackgroundEvents.Size(), ShouldBeZeroValue)
		em.RunEventLoop()
		e := NewEvent(f)
		em.AppendEvent(e)
		Convey("test output event", func() {
			ev := NewEvent(outputFn)
			output := &OutputDTO{}
			em.AppendEvent(ev)
			<-ev.OnOutputChanged()
			ev.LoadOutput(output)
			So(output.TestFieldOfOutput, ShouldEqual, "Some data")

			// time.Sleep(time.Second * 5)
		})
	})
}
