package main

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

type Event struct {
	uuid   uuid.UUID
	event  func(context.Context, *Event)
	cancel context.CancelFunc
	ctx    context.Context

	mu            sync.Mutex
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

func (e *Event) LoadInput(DTO *interface{}) {
	json.Unmarshal(e.input, DTO)
}

func (e *Event) SetOutput(DTO interface{}) {
	e.LockUnlock(func() {
		e.output, _ = json.Marshal(DTO)
	})
}

func (e *Event) LoadOutput(DTO *interface{}) {
	json.Unmarshal(e.output, DTO)
}

func NewEvent() *Event {
	uuid, _ := uuid.NewUUID()
	boolCh := make(chan bool)
	return &Event{uuid: uuid, outputChanged: boolCh}
}
