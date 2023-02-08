package main

import "context"

type StockmanAPIFunc[T any, K any, B any] func(context.Context, chan T, chan K, B, context.CancelFunc)
type AnyEvent = Event[any, any, any]

/*
	Fully generic class. Can be decleared type of output data and input data.
	It helps always to know which kind of data that event is.
*/
type Event[T any, K any, B any] struct {
	StockmanAPIFunc StockmanAPIFunc[T, K, B] // A service API func to use stockman service
	Output          chan T                   // Output data channel to get actual data from process
	Input           chan K                   // Input data channel to send actual data inside process
	Data            B                        // First data to send with event just
	ctx             context.Context          // Context of funning event
	cancelFunc      context.CancelFunc       // Cancel func of running event
}

func (e *Event[T, K, B]) Run() {
	e.StockmanAPIFunc(e.ctx, e.Output, e.Input, e.Data, e.cancelFunc)
}

func (e *Event[T, K, B]) Done() <-chan struct{} {
	return e.ctx.Done()
}

func (e *Event[T, K, B]) StopEvent() {
	e.cancelFunc()
}

func NewEvent[T any, K any, B any]() *Event[T, K, B] {

	ctxBG := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctxBG)

	output := make(chan T)
	input := make(chan K)

	return &Event[T, K, B]{ctx: ctxWithCancel, cancelFunc: cancelFunc, Output: output, Input: input}
}
