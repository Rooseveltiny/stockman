package stockmanapi

import (
	core "stockman/source/stockman_core"
	sandbox "stockman/source/stockman_sandbox"
)

/*
That's complete test event call.
*/
func TestStockmanAPI() *core.StockmanResponse[sandbox.FirstTestDTO] {

	// init event object
	ev := core.NewEvent(sandbox.FirstServiceTestFn)

	// append it to main event loop
	core.SystemEvents_Manager.AppendEvent(ev)

	// wait untill response is ready
	<-ev.OnOutputChanged()

	// init empty dto and load it
	d := &sandbox.FirstTestDTO{}
	ev.LoadOutput(d)
	stockmanResonse := core.NewStockmanResponse(*d, nil)

	// return response with dto
	return stockmanResonse

}
