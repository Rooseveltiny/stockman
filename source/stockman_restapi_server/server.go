package stockmanrestapiserver

import (
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
	logger.L.Fatalln(ras.httpserver.ListenAndServe())
	logger.L.Info("started rest app server. it is accessed by: localhost:8080")
}

func (ras *RestAPIServer) PutRouter(r *Router) {
	ras.router = r
	ras.httpserver.Handler = r.router
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
