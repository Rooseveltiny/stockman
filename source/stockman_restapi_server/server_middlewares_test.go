package stockmanrestapiserver

import (
	"context"
	"io"
	"net/http"
	"stockman/source/stockman_restapi_server/middlewares"
	"testing"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

/* handle */
func testyHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Header.Get("SomeNewHeader") != "It works!" {
		panic("no header was changed!")
	}
	w.Write([]byte(p[0].Value))
}

/* middleware performers */

func testyMiddlewareChangeHeaders(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Header.Add("SomeNewHeader", "It works!")
}

func testyMiddlewareChangeParam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	p[0].Value = "NewRooseveltiny"
}

func testyMiddlewareChangeParam2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	p[0].Value = p[0].Value + "New"
}

/* middleware test */
var testMiddlewareChangeParam middlewares.Middleware = *middlewares.NewMiddleware(testyMiddlewareChangeParam)
var testMiddlewareChangeBody middlewares.Middleware = *middlewares.NewMiddleware(testyMiddlewareChangeHeaders)
var testMiddlewareChangeParam2 middlewares.Middleware = *middlewares.NewMiddleware(testyMiddlewareChangeParam2)

/* middleware set test */
var testMiddlewareSet middlewares.MiddlewareSet = *middlewares.NewMiddlewareSet(&testMiddlewareChangeParam, &testMiddlewareChangeBody)
var testMiddlewareSet2 middlewares.MiddlewareSet = *middlewares.NewMiddlewareSet(&testMiddlewareChangeParam2)
var testMiddlewareConcat middlewares.MiddlewareSet = *middlewares.ConcatMiddlewareSets(&testMiddlewareSet, &testMiddlewareSet2)

/* handlers */
var testyHand *Hand = NewHand(http.MethodGet, "/test_middleware/:user", testyHandle, &testMiddlewareSet)
var testyHand2 *Hand = NewHand(http.MethodGet, "/test_middleware_concat/:user", testyHandle, &testMiddlewareConcat)

func TestMiddlewareRun(t *testing.T) {
	/* prepare server */
	ctx := context.Background()
	CommonHandlers.AppendHandle(*testyHand)
	CommonHandlers.AppendHandle(*testyHand2)
	closeServer := StartRestAPIServer(ctx)
	defer closeServer()
	Convey("test change param middleware", t, func() {
		resp, err := http.Get("http://localhost:8080/test_middleware/Rooseveltiny")
		So(err, ShouldBeNil)
		text_b, errRead := io.ReadAll(resp.Body)
		So(errRead, ShouldBeNil)
		So(string(text_b), ShouldEqual, "NewRooseveltiny")
		Convey("test concat middleware", func() {
			response, errR := http.Get("http://localhost:8080/test_middleware_concat/Rooseveltiny")
			So(errR, ShouldBeNil)
			text_bin, errReadBody := io.ReadAll(response.Body)
			So(errReadBody, ShouldBeNil)
			text := string(text_bin)
			So(text, ShouldEqual, "NewRooseveltinyNew")
		})
	})
}
