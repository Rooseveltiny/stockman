package main

type EventsPool struct {
	EventsList *EventsList
}

func (ep *EventsPool) AppendEvent(event *AnyEvent) {
	ep.EventsList.AppendEvent(event)
}

func (ep *EventsPool) RetrieveEvent() *AnyEvent {
	return ep.EventsList.RetrieveEvent()
}

func (ep *EventsPool) Size() int {
	return ep.EventsList.Size()
}

func NewEventsPool() *EventsPool {
	eventsList := NewEventsList()
	return &EventsPool{EventsList: eventsList}
}
