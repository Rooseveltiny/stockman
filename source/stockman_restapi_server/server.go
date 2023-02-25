package stockmanrestapiserver

import (
	"context"
	"errors"
	"net/http"
	logger "stockman/source/stockman_logger"
	"time"
)

type RestAPIServer struct {
	router     *Router
	httpserver *http.Server
}

func (ras *RestAPIServer) Start() {
	if ras.router == nil {
		logger.L.Errorln(errors.New("no router given"))
	}
	logger.L.Info("started rest app server. it is accessed by: localhost:8080")
	logger.L.Fatalln(ras.httpserver.ListenAndServe())
}

func (ras *RestAPIServer) PutRouter(r *Router) {
	ras.router = r
	ras.httpserver.Handler = r.router
}

func (ras *RestAPIServer) ShoutdownServer(ctx context.Context) {
	ras.httpserver.Shutdown(ctx)
}

func NewRestAPIServer() *RestAPIServer {
	https := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	nrest := RestAPIServer{
		httpserver: https,
	}
	return &nrest
}

/* init and start rest api server */
func StartRestAPIServer(ctx context.Context) func() {

	/* init server instance */
	restApiServer := NewRestAPIServer()

	/* init router for server and register all routes collections */
	router := NewRouter()
	router.RegisterAllRoutes()

	/* give server routes */
	restApiServer.PutRouter(router)

	/* start server listening */
	go restApiServer.Start()

	ShoutdownFunc := func() { restApiServer.ShoutdownServer(ctx) }
	return ShoutdownFunc
}
