package main

import "context"

type StockmanAPIFunc[T any, K any, B any] func(context.Context, T, K, B)
type AnyEvent Event[any, any, any]

/*
	Fully generic class. Can be decleared type of output data and input data.
	It helps always to know which kind of data that event is.
*/
type Event[T any, K any, B any] struct {
	StockmanAPIFunc StockmanAPIFunc[T, K, B]
	Output          chan<- T // Output data channel to get actual data from process
	Input           <-chan K // Input data channel to send actual data inside process
	Data            B        // First data to send with event just
}

func (e *Event[T, K, B]) Run() {}

func NewEvent[T any, K any, B any]() *Event[T, K, B] {
	return &Event[T, K, B]{}
}
