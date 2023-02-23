package vcpostgres

import (
	"context"
	"stockman/source/stockman_dbs/client/postgresql"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	postgresutils "stockman/source/stockman_dbs/postgres_utils"
	logger "stockman/source/stockman_logger"
	"testing"

	"github.com/google/uuid"
	"github.com/smartystreets/goconvey/convey"
)

func TestRepository(t *testing.T) {
	ctx := context.TODO()
	client, _ := postgresql.NewClient(ctx, *postgresql.NewTestConfig())

	/* before test */
	postgresutils.PrepareTestPostgresSQL(ctx, client)
	/* after test */
	defer postgresutils.DropPreparedTestPostgresSQL(ctx, client)

	convey.Convey("init new repository", t, func() {
		repo := NewRepository(ctx, logger.L)
		convey.So(repo.client, convey.ShouldNotBeNil)
		convey.So(repo.logger, convey.ShouldNotBeNil)
		convey.Convey("create video camera", func() {
			videocameraDTO := videocamera.CameraCreateDTO{
				Address:  "testaddress",
				Port:     "5454",
				Login:    "stockman",
				Password: "43598340f345erg$T#$R#",
			}
			link, err := repo.Create(ctx, videocameraDTO)
			convey.So(err, convey.ShouldBeNil)
			_, parseErr := uuid.Parse(link)
			convey.So(parseErr, convey.ShouldBeNil)
			convey.Convey("test get this row from database", func() {
				f, _ := repo.GetByLink(ctx, link)
				convey.So(f.Address, convey.ShouldEqual, "testaddress")
				convey.So(f.Port, convey.ShouldEqual, "5454")
				convey.So(f.Login, convey.ShouldEqual, "stockman")
				convey.So(f.Password, convey.ShouldEqual, "43598340f345erg$T#$R#")
			})
			convey.Convey("test get all from db", func() {
				videocameraDTO := videocamera.CameraCreateDTO{
					Address:  "testaddress",
					Port:     "5454",
					Login:    "stockman",
					Password: "43598340f345erg$T#$R#",
				}
				repo.Create(ctx, videocameraDTO)
				repo.Create(ctx, videocameraDTO)
				repo.Create(ctx, videocameraDTO)
				repo.Create(ctx, videocameraDTO)
				repo.Create(ctx, videocameraDTO)
				vcddtos, errgetAll := repo.GetAll(ctx)
				convey.So(errgetAll, convey.ShouldBeNil)
				convey.So(len(vcddtos), convey.ShouldEqual, 7 /* one more because first func called twice */)
			})
		})
	})
}
