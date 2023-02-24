package core

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

type EventFn = func(context.Context, *Event)

type Event struct {
	uuid   uuid.UUID
	event  EventFn
	cancel context.CancelFunc
	ctx    context.Context

	mu            sync.Mutex
	err           error
	input         []byte
	outputChanged chan bool
	output        []byte
}

func (e *Event) RunEvent() {
	go e.event(e.ctx, e)
}

func (e *Event) CancelEvent() {
	e.cancel()
}

func (e *Event) NotifyOutputChanged() {
	e.outputChanged <- true
}

func (e *Event) OnOutputChanged() <-chan bool {
	return e.outputChanged
}

func (e *Event) LockUnlock(f func()) {
	e.mu.Lock()
	defer e.mu.Unlock()
	f()
}

func (e *Event) SetInput(DTO interface{}) {
	e.LockUnlock(func() {
		e.input, _ = json.Marshal(DTO)
	})
}

func (e *Event) LoadInput(DTO interface{}) {
	json.Unmarshal(e.input, DTO)
}

func (e *Event) SetOutput(DTO interface{}) {
	e.LockUnlock(func() {
		e.output, _ = json.Marshal(DTO)
	})
}

func (e *Event) LoadOutput(DTO interface{}) {
	json.Unmarshal(e.output, DTO)
}

func (e *Event) SetError(err error) {
	e.LockUnlock(func() {
		e.err = err
	})
}

func (e *Event) Error() error {
	return e.err
}

func NewEvent(fn EventFn) *Event {
	uuid, _ := uuid.NewUUID()
	boolCh := make(chan bool)
	ctx := context.Background()
	cancelCtx, cancelFn := context.WithCancel(ctx)

	return &Event{
		uuid:          uuid,
		outputChanged: boolCh,
		cancel:        cancelFn,
		ctx:           cancelCtx,
		event:         fn,
	}
}
