package main

import "context"

type SystemEventsManager struct {
	Events           *EventsPool
	BackgroundEvents *EventsPool
	NewEventFlag     chan bool
	ctx              context.Context
}

func (sem *SystemEventsManager) RunEventLoop() {
	for {
		select {
		case <-sem.ctx.Done():
			return
		case <-sem.NewEventFlag: // Run sync events
			sem.RunEvent()
		}
		// run bg events here
	}
}

func (sem *SystemEventsManager) AppendEvent(event *Event) {
	sem.Events.AppendEvent(event)
	sem.NewEventFlag <- true
}

func (sem *SystemEventsManager) RunEvent() {
	if sem.Events.Size() != 0 {
		event := sem.Events.PopEvent()
		event.RunEvent()
	}
}

func (sem *SystemEventsManager) AppendBackgroundEvent(event *Event) {
	sem.BackgroundEvents.AppendEvent(event)
}

func (sem *SystemEventsManager) RunBackgroundEvent() {
	if sem.Events.Size() != 0 {
		event := sem.BackgroundEvents.PopEvent()
		event.RunEvent()
	}
}

func NewSystemEventsManager() *SystemEventsManager {
	evPool := NewEventsPool()
	evBGPool := NewEventsPool()
	return &SystemEventsManager{Events: evPool, BackgroundEvents: evBGPool}
}
