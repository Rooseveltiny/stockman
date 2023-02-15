package stockmanapi

import (
	core "stockman/source/stockman_core"
	sandbox "stockman/source/stockman_sandbox"
)

/*
That's complete test event call.
*/
func TestStockmanAPI() *core.StockmanResponse[sandbox.FirstTestDTO] {
	ev := core.NewEvent(sandbox.FirstServiceTestFn)
	core.SystemEvents_Manager.AppendEvent(ev)
	<-ev.OnOutputChanged()
	d := &sandbox.FirstTestDTO{}
	ev.LoadOutput(d)
	stockmanResonse := core.NewStockmanResponse(*d, nil)
	return stockmanResonse
}
