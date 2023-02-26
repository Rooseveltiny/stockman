package stockmanrestapiserver

import (
	"stockman/source/stockman_restapi_server/middlewares"

	"github.com/julienschmidt/httprouter"
)

/*
struct which has description of endpoint of a system
- method (post, put, delete, get e.c.)
- path (url where this endpoint waits for request)
- handle - function which run the given request
*/
type Hand struct {
	Method        string
	Path          string
	Handle        httprouter.Handle
	middlewareSet *middlewares.MiddlewareSet
}

/* if theres no middlewares was set then give to route just handle else give wrapped with middlewares handle */
func (h *Hand) GetHandle() httprouter.Handle {
	if h.middlewareSet == nil {
		return h.Handle
	} else {
		return h.middlewareSet.MiddlewareWrap(h.Handle)
	}
}

func NewHand(method, path string, handle httprouter.Handle, m *middlewares.MiddlewareSet) *Hand {
	h := &Hand{
		Method: method,
		Path:   path,
		Handle: handle,
	}
	if m != nil {
		h.middlewareSet = m
	}
	return h
}

type RoutesCollection struct {
	handlers []Hand
}

func (rc *RoutesCollection) LoadRouterWithRoutes(r *Router) {
	for _, h := range rc.handlers {
		r.AddHandle(h)
	}
}

func (rc *RoutesCollection) AppendHandle(h Hand) {
	rc.handlers = append(rc.handlers, h)
}

func (rc *RoutesCollection) Handlers() []Hand {
	return rc.handlers
}

func NewRoutesCollection() *RoutesCollection {
	return &RoutesCollection{}
}
