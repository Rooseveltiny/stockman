package main

type EventsPool struct {
	Events *EventsList
}

func (ep *EventsPool) Size() int {
	return ep.Events.Size()
}

func (ep *EventsPool) AppendEvent(event *Event) {
	ep.Events.AppendEvent(event)
}

func (ep *EventsPool) PopEvent() *Event {
	return ep.Events.PopEvent()
}

func NewEventsPool() *EventsPool {
	eventsList := NewEventsList()
	return &EventsPool{Events: eventsList}
}
