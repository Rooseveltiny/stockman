package main

type EventsPool struct {
	EventsList *EventsList
}

func (ep *EventsPool) RetrieveEvent() *AnyEvent {
	return ep.EventsList.RetrieveEvent()
}

func NewEventsPool() *EventsPool {
	eventsList := NewEventsList()
	return &EventsPool{EventsList: eventsList}
}
