package main

type EventsList struct {
	Events []*AnyEvent
}

func (el *EventsList) AppendEvent(event *AnyEvent) {
	el.Events = append(el.Events, event)
}

func (el *EventsList) RetrieveEvent() *AnyEvent {
	var retrievedEvent AnyEvent
	retrievedEvent, el.Events = *el.Events[len(el.Events)-1], el.Events[:len(el.Events)-1]
	return &retrievedEvent
}

func (el *EventsList) Size() int {
	return len(el.Events)
}

func (el *EventsList) ClearEventsList() {
	for len(el.Events) != 0 {
		el.RetrieveEvent()
	}
}

func NewEventsList() *EventsList {
	return &EventsList{}
}
