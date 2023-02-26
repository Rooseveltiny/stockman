package middlewares

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func someTestHandleTest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func TestMiddleware(t *testing.T) {
	Convey("test init middleware instance", t, func() {
		m := NewMiddleware(someTestHandleTest)
		So(m.middleware, ShouldNotBeNil)
		h := m.GetHandle()
		So(h, ShouldEqual, someTestHandleTest)
		Convey("test middlewareset init", func() {
			ms := NewMiddlewareSet(m)
			So(len(ms.middlewares), ShouldEqual, 1)
			Convey("test init middlewareset with some middlewares", func() {
				m1 := NewMiddleware(someTestHandleTest)
				m2 := NewMiddleware(someTestHandleTest)
				m3 := NewMiddleware(someTestHandleTest)
				msWithSomeMiddlewares := NewMiddlewareSet(m1, m2, m3)
				So(len(msWithSomeMiddlewares.middlewares), ShouldEqual, 3)
				Convey("test middleware concationation sets", func() {
					newMS := ConcatMiddlewareSets(ms, msWithSomeMiddlewares)
					So(len(newMS.middlewares), ShouldEqual, 4)
					Convey("test middleware set append func", func() {
						newMS.AppendMiddleware(m1)
						So(len(newMS.middlewares), ShouldEqual, 5)
					})
				})
			})
		})
	})
}
