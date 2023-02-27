package middlewares

import (
	"fmt"
	"net/http"
	logger "stockman/source/stockman_logger"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	middleware httprouter.Handle
}

func (m *Middleware) GetHandle() httprouter.Handle {
	return m.middleware
}

func NewMiddleware(h httprouter.Handle) *Middleware {
	return &Middleware{
		middleware: h,
	}
}

type MiddlewareSet struct {
	middlewares []*Middleware
}

func (ms *MiddlewareSet) AppendMiddleware(m *Middleware) {
	ms.middlewares = append(ms.middlewares, m)
}

/* new function wich performs all middleware funcs and gives http.Handle at last to router */
func (ms *MiddlewareSet) MiddlewareWrap(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		/* log gotten request */
		logger.L.Infoln(fmt.Sprintf("method %s", r.Method))

		for _, m := range ms.middlewares {
			/* calls all middlewares before calling main handle func */
			m.middleware(w, r, p)
		}
		/* call handle func at last */
		h(w, r, p)
	}
}

func NewMiddlewareSet(h ...*Middleware) *MiddlewareSet {
	newMiddlewareSet := &MiddlewareSet{}
	for _, m := range h {
		newMiddlewareSet.AppendMiddleware(m)
	}
	return newMiddlewareSet
}

/* joins several middleware sets */
func ConcatMiddlewareSets(ms ...*MiddlewareSet) *MiddlewareSet {
	resultMiddlewareSet := NewMiddlewareSet()
	for _, middleWareSet := range ms {
		for _, middleware := range middleWareSet.middlewares {
			resultMiddlewareSet.AppendMiddleware(middleware)
		}
	}
	return resultMiddlewareSet
}
