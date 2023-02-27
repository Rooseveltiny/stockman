package apitest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	videocamera "stockman/source/stockman_dbs/data_models/video_camera"
	"testing"

	"github.com/google/uuid"
	"github.com/smartystreets/goconvey/convey"
)

const videomanagerAPIPath string = "/api/v1/video_camera"

func TestVideomanagerAPI(t *testing.T) {
	ctx := context.Background()
	runTest(ctx, func(ctx context.Context) {
		convey.Convey("test create video camera", t, func() {
			d := videocamera.CameraCreateDTO{
				Address:  "http://localhost:9090",
				Port:     "5432",
				Login:    "Roosy",
				Password: "@#^#$T#%$RGER",
			}
			body, err := json.Marshal(d)
			fmt.Println(err)
			resp, errResp := http.Post(testBaseURL+videomanagerAPIPath+"/create_camera", textPlain, bytes.NewBuffer(body))
			convey.So(errResp, convey.ShouldBeNil)
			defer resp.Body.Close()
			bodyBytes, _ := io.ReadAll(resp.Body)
			_, parseErr := uuid.Parse(string(bodyBytes))
			convey.So(parseErr, convey.ShouldBeNil)
		})
	})
}
