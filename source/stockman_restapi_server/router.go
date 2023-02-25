package stockmanrestapiserver

import "github.com/julienschmidt/httprouter"

type Router struct {
	router *httprouter.Router
}

func (r *Router) AddHandle(h Hand) {
	r.router.Handle(h.Method, h.Path, h.Handle)
}

func (r *Router) RegisterAllRoutes() {
	/* init all handlers of a system here */
	AllHandlers.LoadRouterWithRoutes(r)
	/* add new here ... */
}

func NewRouter() *Router {
	newhttprouter := httprouter.New()
	router := Router{
		router: newhttprouter,
	}
	return &router
}

/* routes collection */
var AllHandlers *RoutesCollection = NewRoutesCollection()
