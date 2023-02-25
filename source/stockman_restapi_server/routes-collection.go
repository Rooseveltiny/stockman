package stockmanrestapiserver

import (
	"github.com/julienschmidt/httprouter"
)

type Hand struct {
	Method string
	Path   string
	Handle httprouter.Handle
}

func NewHand(method, path string, handle httprouter.Handle) *Hand {
	return &Hand{
		Method: method,
		Path:   path,
		Handle: handle,
	}
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
