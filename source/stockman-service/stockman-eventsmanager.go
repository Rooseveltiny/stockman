package main

type SystemEventsManager struct {
	Events *EventsPool

	BackgroundEvents        *EventsPool
	RunningBackgroundEvents *EventsPool
}

func (sem *SystemEventsManager) RunEventLoop() {
	for {
		for sem.Events.Size() != 0 {
			event := sem.Events.RetrieveEvent()
			go sem.RunEvent(event)
		}
		for sem.BackgroundEvents.Size() != 0 {
			backgroundEvent := sem.BackgroundEvents.RetrieveEvent()
			go sem.RunBackgroundEvent(backgroundEvent)
		}
	}
}

func (sem *SystemEventsManager) AppendEvent(event *AnyEvent) {
	sem.Events.AppendEvent(event)
}

func (sem *SystemEventsManager) RunEvent(event *AnyEvent) {
	event.Run()
}

func (sem *SystemEventsManager) AppendBackgroundEvent(event *AnyEvent) {
	sem.BackgroundEvents.AppendEvent(event)
}

func (sem *SystemEventsManager) RunBackgroundEvent(event *AnyEvent) {
	sem.RunningBackgroundEvents.AppendEvent(event) // put event into running pool
}

func NewSystemEventsManager() *SystemEventsManager {
	eventsPool := NewEventsPool()
	bgEventsPool := NewEventsPool()
	return &SystemEventsManager{Events: eventsPool, BackgroundEvents: bgEventsPool}
}
