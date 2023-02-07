package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/smartystreets/goconvey/convey"
)

/*
	Testing EVENT
*/
func TestEvent(t *testing.T) {
	event := NewEvent(0, nil)
	convey.Convey("Test event init", t, func() {
		convey.So(event, convey.ShouldNotBeNil)
		convey.So(event.Input, convey.ShouldEqual, 0)
		convey.Convey("Validate generated uuid", func() {
			_, err := uuid.Parse(event.uuid.String())
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

/*
	Testing Events Pool
*/
func TestEventsList(t *testing.T) {
	eventList := NewEventsList()
	event1 := NewEvent(0, nil)
	// uuid1 := event1.uuid
	event2 := NewEvent(0, nil)
	uuid2 := event2.uuid
	event3 := NewEvent(0, nil)
	// uuid3 := event3.uuid

	convey.Convey("Test event list init", t, func() {
		convey.So(eventList.Events, convey.ShouldBeEmpty)
		convey.So(eventList.Size(), convey.ShouldBeZeroValue)
		convey.Convey("Test adding new event", func() {
			eventList.AppendEvent(event1)
			eventList.AppendEvent(event2)
			eventList.AppendEvent(event3)
			convey.So(eventList.Events[0], convey.ShouldEqual, event1)
			convey.So(eventList.Size(), convey.ShouldEqual, 3)
			convey.Convey("Test popping and element", func() {
				popedElement := eventList.PopEvent()
				convey.So(popedElement, convey.ShouldEqual, event1)
				eventList.AppendEvent(popedElement)
				convey.So(eventList.Size(), convey.ShouldEqual, 3)
				popedElementByUUID := eventList.PopEventByUUID(uuid2)
				convey.So(popedElementByUUID, convey.ShouldEqual, event2)
				convey.Convey("Test clearing events list", func() {
					eventList.ClearList()
					convey.So(eventList.Size(), convey.ShouldEqual, 0)
				})
			})
		})
	})
}

/*
	Testing Stockman main component
*/
func TestEventsPool(t *testing.T) {
	convey.Convey("Test events pool init", t, func() {
		eventsPool := NewEventsPool()
		convey.So(eventsPool.Events, convey.ShouldNotBeNil)
		convey.So(eventsPool.BackgroundEvents, convey.ShouldNotBeNil)
		convey.So(eventsPool.RunningBackgroundEvents, convey.ShouldNotBeNil)
		convey.Convey("Testing push event", func() {
			event := NewEvent(0, nil)
			eventsPool.PushEvent(event)
			convey.So(eventsPool.Events.Size(), convey.ShouldEqual, 1)
			time.Sleep(time.Millisecond * 500)
			convey.Convey("Test new pushed event should be runned and removed from pool", func() {
				convey.So(eventsPool.Events.Size(), convey.ShouldEqual, 0)
			})
			time.Sleep(time.Second * 1)
		})
	})
}
