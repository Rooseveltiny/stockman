package stockmanapi

import (
	core "stockman/source/stockman_core"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFirstAPICall(t *testing.T) {

	// init main stockman component
	ss := core.NewStockmanService()
	ss.StartSystem()

	convey.Convey("test first api call", t, func() {
		r := TestStockmanAPI()
		convey.So(r.DTOData.TestFieldDTO, convey.ShouldEqual, "Hello services!")
	})

}
