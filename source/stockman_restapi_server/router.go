package stockmanrestapiserver

import "github.com/julienschmidt/httprouter"

type Router struct {
	router *httprouter.Router
}

func (r *Router) RegisterAllRoutes() {

}

func NewRouter() *Router {
	newhttprouter := httprouter.New()
	router := Router{
		router: newhttprouter,
	}
	return &router
}
