package core

type EventsList struct {
	Events []*Event
}

func (el *EventsList) AppendEvent(event *Event) {
	el.Events = append(el.Events, event)
}

func (el *EventsList) PopEvent() *Event {
	var event *Event
	event, el.Events = el.Events[len(el.Events)-1], el.Events[:len(el.Events)-1]
	return event
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
