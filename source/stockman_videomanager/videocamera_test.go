package stockmanvideomanager

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	postgresutils "stockman/source/stockman_dbs/postgres_utils"
	"testing"

	"github.com/google/uuid"
	convey "github.com/smartystreets/goconvey/convey"
)

func TestVideoCamera(t *testing.T) {
	ctx := context.TODO()
	cl, _ := postgresql.NewTestClient(ctx)
	postgresutils.PrepareTestPostgresSQL(ctx, cl)

	convey.Convey("create new camera", t, func() {
		cdto := videocamera.CameraCreateDTO{
			Address:  "some address",
			Port:     "4545",
			Login:    "Saveliy",
			Password: "09534059",
		}
		link, err := CreateNewCamera(ctx, cdto)
		convey.So(err, convey.ShouldBeNil)
		_, errParse := uuid.Parse(link)
		convey.So(errParse, convey.ShouldBeNil)

		convey.Convey("get video camera entity", func() {
			vc, err := GetVideoCamera(ctx, link)
			convey.So(err, convey.ShouldBeNil)
			convey.So(vc.Address, convey.ShouldEqual, cdto.Address)
			convey.So(vc.Login, convey.ShouldEqual, cdto.Login)
			convey.So(vc.Password, convey.ShouldEqual, cdto.Password)
			convey.So(vc.Port, convey.ShouldEqual, cdto.Port)
		})
	})

	postgresutils.DropPreparedTestPostgresSQL(ctx, cl)
}
