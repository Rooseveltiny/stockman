package main

type SystemEventsManager struct {
	Events           *EventsPool
	BackgroundEvents *EventsPool
}

func (sem *SystemEventsManager) RunEventLoop() {
	// running infinite loop to run all ongoing events
}

func AppendEvent(*AnyEvent) {}

func AppendBackgroundEvent(*AnyEvent) {}

func NewSystemEventsManager() *SystemEventsManager {
	eventsPool := NewEventsPool()
	bgEventsPool := NewEventsPool()
	return &SystemEventsManager{Events: eventsPool, BackgroundEvents: bgEventsPool}
}
