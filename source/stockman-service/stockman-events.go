package main

import (
	"context"

	"github.com/google/uuid"
)

/*
	function to perform in system
*/
type EventFunc func(ctx context.Context, input interface{}, output chan interface{})

/*
	Events pool is a special collection to rule all the events of a system
*/
type EventsPool struct {
	Events                  *EventsList
	BackgroundEvents        *EventsList
	RunningBackgroundEvents *EventsList
}

func (ep *EventsPool) RunEvents() {
	for {
		for ep.Events.Size() != 0 {
			eventToPerfrom := ep.PopEvent()
			ep.RunEvent(eventToPerfrom)
		}
		for ep.BackgroundEvents.Size() != 0 {
			backgroundEventToPerform := ep.PopBackgroundEvent()
			ep.RunBackgroundEvent(backgroundEventToPerform)
			ep.appendRunningBackgroundEvent(backgroundEventToPerform)
		}
	}
}

func (ep *EventsPool) ClearEndedBackgroundEvents() {
	for {
		for ep.RunningBackgroundEvents.Size() != 0 {
			for _, event := range ep.RunningBackgroundEvents.Events {
				<-event.ContextWithCancel.Done()
				ep.PopRunningBackgroundEventByUUID(event.uuid)
			}
		}
	}
}

func (ep *EventsPool) PopEvent() *Event {
	return ep.Events.PopEvent()
}

func (ep *EventsPool) PopBackgroundEvent() *Event {
	return ep.BackgroundEvents.PopEvent()
}

func (ep *EventsPool) PopRunningBackgroundEventByUUID(uuid uuid.UUID) *Event {
	return ep.RunningBackgroundEvents.PopEventByUUID(uuid)
}

func (ep *EventsPool) PushEvent(event *Event) {
	ep.Events.AppendEvent(event)
}

func (ep *EventsPool) PushBackgroundEvent(event *Event) {
	ep.BackgroundEvents.AppendEvent(event)
}

func (ep *EventsPool) appendRunningBackgroundEvent(event *Event) {
	ep.RunningBackgroundEvents.AppendEvent(event)
}

func (ep *EventsPool) RunEvent(event *Event) {}

func (ep *EventsPool) RunBackgroundEvent(event *Event) {}

func (ep *EventsPool) StopRunningBackgroundEvent(uuid uuid.UUID) {
	runningBackgroundEvent := ep.PopRunningBackgroundEventByUUID(uuid)
	runningBackgroundEvent.StopEvent()
}

func NewEventsPool() *EventsPool {
	Events := NewEventsList()
	BackgroundEvents := NewEventsList()
	RunningBackgroundEvents := NewEventsList()
	eventsPool := EventsPool{Events, BackgroundEvents, RunningBackgroundEvents}
	go eventsPool.RunEvents()
	return &eventsPool
}

/*
	Events list is a special collection structure data which helps to
	manipulate with events providng useful methods
*/
type EventsList struct {
	Events []*Event
}

func (el *EventsList) AppendEvent(event *Event) {
	el.Events = append(el.Events, event)
}

func (el *EventsList) PopEvent() *Event {
	var event *Event
	event, el.Events = el.Events[0], el.Events[1:]
	return event
}
func (el *EventsList) PopEventByUUID(uuid uuid.UUID) *Event {
	for i, v := range el.Events {
		if v.uuid == uuid {
			el.Events = append(el.Events[:i], el.Events[i+1:]...)
			return v
		}
	}
	return nil
}

func (el *EventsList) Size() int {
	return len(el.Events)
}

func (el *EventsList) ClearList() {
	for el.Size() != 0 {
		el.PopEvent()
	}
}

func NewEventsList() *EventsList {
	return &EventsList{}
}

/*
	Event is a special entity which performs functions of a system
*/
type Event struct {
	Output            chan interface{}
	Input             interface{}
	ContextWithCancel context.Context
	EventFunc         EventFunc

	uuid      uuid.UUID
	stopEvent context.CancelFunc
}

func (e *Event) ExecuteEvent() {
	go e.EventFunc(e.ContextWithCancel, e.Input, e.Output)
}

func (e *Event) StopEvent() {
	e.stopEvent()
}

func NewEvent(inputData interface{}, eventFunc EventFunc) *Event {

	uuid := uuid.New()
	contextWithCancel, cancelFunc := context.WithCancel(context.Background())
	var output = make(chan interface{})

	return &Event{
		Input:             inputData,
		EventFunc:         eventFunc,
		uuid:              uuid,
		ContextWithCancel: contextWithCancel,
		stopEvent:         cancelFunc,
		Output:            output,
	}
}
