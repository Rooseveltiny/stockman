package stockmanapi

import (
	core "stockman/source/stockman_core"
	sandbox "stockman/source/stockman_sandbox"
)

func TestStockmanAPI() *core.StockmanResponse[string] {
	ev := core.NewEvent(sandbox.FirstServiceTestFn)
	core.SystemEvents_Manager.AppendEvent(ev)

	return nil
}
