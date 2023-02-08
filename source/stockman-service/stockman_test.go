package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEvent(t *testing.T) {
	Convey("test event init", t, func() {
		event := NewEvent[string, string]("")
		So(event.ctx, ShouldNotBeNil)
		So(event.Data, ShouldEqual, "")
		So(event.Input, ShouldNotBeNil)
		So(event.Output, ShouldNotBeNil)
		Convey("test event stop", func() {
			func() {
				go func() {
					time.Sleep(time.Millisecond * 500)
					event.StopEvent()
				}()
				for {
					<-event.Done()
					break
				}
			}()
		})
	})
}

func TestEventList(t *testing.T) {
	Convey("test eventlist init", t, func() {
		eventList := NewEventsList()
		So(eventList.Events, ShouldBeEmpty)
		So(eventList.Size(), ShouldBeZeroValue)
		Convey("test eventlist append event", func() {
			event := NewEvent[string, string, string]("")
			eventList.AppendEvent(event)
		})
	})
}
