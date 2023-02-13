package main

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type Event struct {
	uuid      uuid.UUID
	eventFunc func(context.Context, *Event)

	ctx          context.Context
	cancelFunc   context.CancelFunc
	onPause      chan bool
	inputStream  chan []byte
	outputStream chan []byte
	input        []byte
	output       []byte
}

func (e *Event) Run() {
	e.eventFunc(e.ctx, e)
}

func (e *Event) Play() {
	e.onPause <- false
}

func (e *Event) Pause() {
	e.onPause <- true
}

func (e *Event) FinishEvent() {
	e.cancelFunc()
}

func (e *Event) EventDone() <-chan struct{} {
	return e.ctx.Done()
}

func (e *Event) LoadLastOutputValue(DTO *interface{}) {
	json.Unmarshal(e.output, DTO)
}

func (e *Event) LoadLastInputValue(DTO *interface{}) {
	json.Unmarshal(e.input, DTO)
}

func (e *Event) SetInputStreamValue(DTO interface{}) {
	inputValue, _ := json.Marshal(DTO)
	e.input = inputValue
	e.inputStream <- inputValue
}

func (e *Event) SetOutputStreamValue(DTO interface{}) {
	outputValue, _ := json.Marshal(DTO)
	e.output = outputValue
	e.outputStream <- outputValue
}

func NewEvent() *Event {
	ctx := context.Background()
	ctxCancel, cancelFunc := context.WithCancel(ctx)

	emptyInputStream := make(chan []byte)
	emptyOutputStream := make(chan []byte)
	onPause := make(chan bool)
	onPause <- false

	uuid, _ := uuid.NewUUID()

	e := &Event{
		uuid: uuid,

		ctx:        ctxCancel,
		cancelFunc: cancelFunc,

		onPause: onPause,

		inputStream:  emptyInputStream,
		outputStream: emptyOutputStream,
	}

	return e
}
