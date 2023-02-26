package stockmanrestapiserver

import (
	"context"
	"io/ioutil"
	"net/http"
	core "stockman/source/stockman_core"
	"stockman/source/stockman_dbs/client/postgresql"
	postgresutils "stockman/source/stockman_dbs/postgres_utils"
	logger "stockman/source/stockman_logger"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/smartystreets/goconvey/convey"
)

const Test_Host = "http://localhost:8080"
const Test_URL = "/test_api"
const Test_Message = "Hello Restapi!"

/* handle funcs */
func HandlerFuncForTest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte(Test_Message))
}

var TestHandler *Hand = NewHand(http.MethodGet, Test_URL, HandlerFuncForTest)

func TestRouter(t *testing.T) {
	convey.Convey("test init router", t, func() {
		router := NewRouter()
		convey.So(router.router, convey.ShouldNotBeNil)
		router.RegisterAllRoutes()
		convey.So(router.router, convey.ShouldNotBeEmpty)
	})
}

func TestRoutesCollection(t *testing.T) {
	convey.Convey("test init routes collections", t, func() {
		routes := NewRoutesCollection()
		convey.So(routes.handlers, convey.ShouldBeNil)
		convey.Convey("test init new hand", func() {
			newHand := NewHand(http.MethodGet, "some_route", HandlerFuncForTest)
			convey.So(newHand.Method, convey.ShouldEqual, http.MethodGet)
			convey.So(newHand.Path, convey.ShouldEqual, "some_route")
			convey.So(newHand.Handle, convey.ShouldEqual, HandlerFuncForTest)
			routes.AppendHandle(*newHand)
			convey.So(routes.handlers, convey.ShouldNotBeEmpty)
		})
	})
}

func TestRestAPIServer(t *testing.T) {
	convey.Convey("test init rest api server", t, func() {
		restapi := NewRestAPIServer()
		router := NewRouter()
		convey.So(restapi.router, convey.ShouldBeNil)
		restapi.PutRouter(router)
		convey.So(restapi.router, convey.ShouldNotBeNil)
		ctx := context.Background()
		go func() {
			logger.L.Info("wait for 1 seconds")
			time.Sleep(time.Second * 1)
			restapi.ShoutdownServer(ctx)
		}()
		restapi.Start()
	})
}

func TestCallRestAPI(t *testing.T) {
	ctx := context.TODO()
	c, errConn := postgresql.GetPostgresClient(ctx)

	AllHandlers.AppendHandle(*TestHandler)

	closeServer := StartRestAPIServer(ctx)
	defer closeServer()
	core.InitAndRunStockmanService()
	postgresutils.PrepareTestPostgresSQL(ctx, c)
	convey.Convey("test api call", t, func() {
		convey.So(errConn, convey.ShouldBeNil)
		resp, err := http.Get(Test_Host + Test_URL)
		convey.So(err, convey.ShouldBeNil)
		text_b, errRead := ioutil.ReadAll(resp.Body)
		convey.So(errRead, convey.ShouldBeNil)
		text := string(text_b)
		convey.So(text, convey.ShouldEqual, Test_Message)
	})
	postgresutils.DropPreparedTestPostgresSQL(ctx, c)
}
