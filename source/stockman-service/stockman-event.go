package main

import "context"

type StockmanAPIFunc[T any, K any, B any] func(context.Context, T, K, B)
type AnyEvent = Event[any, any, any]

/*
	Fully generic class. Can be decleared type of output data and input data.
	It helps always to know which kind of data that event is.
*/
type Event[T any, K any, B any] struct {
	StockmanAPIFunc StockmanAPIFunc[T, K, B]
	Output          chan<- T // Output data channel to get actual data from process
	Input           <-chan K // Input data channel to send actual data inside process
	Data            B        // First data to send with event just
	ctx             context.Context
	cancelFunc      context.CancelFunc
}

func (e *Event[T, K, B]) Run() {

}

func (e *Event[T, K, B]) Done() <-chan struct{} {
	return e.ctx.Done()
}

func (e *Event[T, K, B]) StopEvent() {
	e.cancelFunc()
}

func NewEvent[T any, K any, B any]() *Event[T, K, B] {
	return &Event[T, K, B]{}
}
