package stockmanapi

import (
	core "stockman/source/stockman_core"
	sandbox "stockman/source/stockman_sandbox"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

/*
That's complete test event call.
*/
func StockmanAPITestFunc() *core.StockmanResponse[sandbox.FirstTestDTO] {

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

func TestFirstAPICall(t *testing.T) {
	// init main stockman component
	ss := core.NewStockmanService()
	ss.StartSystem()
	convey.Convey("test first api call", t, func() {
		r := StockmanAPITestFunc()
		convey.So(r.DTOData.TestFieldDTO, convey.ShouldEqual, "Hello services!")
	})
}
