package stockmanapi

import (
	"context"
	core "stockman/source/stockman_core"
	"stockman/source/stockman_dbs/client/postgresql"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	postgresutils "stockman/source/stockman_dbs/postgres_utils"
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

func TestAddVideoCamera(t *testing.T) {
	ctx := context.TODO()
	c, _ := postgresql.GetPostgresClient(ctx)
	postgresutils.PrepareTestPostgresSQL(ctx, c)
	defer postgresutils.DropPreparedTestPostgresSQL(ctx, c)

	ss := core.NewStockmanService()
	ss.StartSystem()

	convey.Convey("test add new camera", t, func() {
		newCamera := videocamera.CameraCreateDTO{
			Address:  "localhost://127.0.0.1",
			Port:     "5432",
			Login:    "Saveliy Trif",
			Password: "manyAprilsPastManyDays32432%#$#$%^#$^@#$#(%(GK))",
		}
		cameraLink := AddVideoCamera(newCamera)
		convey.So(cameraLink.Err, convey.ShouldBeNil)
	})
}
