package stockmanrestapiserver

import "github.com/julienschmidt/httprouter"

type Router struct {
	router *httprouter.Router
}

func (r *Router) AddHandle(h Hand) {
	r.router.Handle(h.Method, h.Path, h.GetHandle())
}

func (r *Router) RegisterAllRoutes() {
	/* init common handlers */
	CommonHandlers.LoadRouterWithRoutes(r)

	/* init all other handlers collections */
	VideoCameraHandlers_V1.LoadRouterWithRoutes(r)
}

func NewRouter() *Router {
	newhttprouter := httprouter.New()
	router := Router{
		router: newhttprouter,
	}
	return &router
}

/* routes common collection. here can be put different common handlers. i.e. test handlers */
var CommonHandlers *RoutesCollection = NewRoutesCollection("")
