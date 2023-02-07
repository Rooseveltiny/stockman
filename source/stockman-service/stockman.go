package main

/*
	Core entity of a system where all events are sent
*/
type Stockman struct {
	EventsPool *EventsPool
}

func NewStockman() *Stockman {
	eventsPool := NewEventsPool()
	return &Stockman{EventsPool: eventsPool}
}
