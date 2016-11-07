package controller

import (
	"fmt"
	"github.com/gauravtayal0/sysutil/services"
	"github.com/gorilla/mux"
	"net/http"
)

var endpoints []string

type WebServer struct {
	Channel <-chan string // channel for incoming communication with controlling infra
	Port    int
}

func NewWebServer(channel <-chan string, port int) *WebServer {
	return &WebServer{Channel: channel, Port: port}
}

func (this *WebServer) Start() error {
	router := NewRouter()
	addr := fmt.Sprintf(":%v", this.Port)

	err := http.ListenAndServe(addr, router)
	return err
}

func NewRouter() *mux.Router {
	endpoints = make([]string, len(routes))
	router := mux.NewRouter().StrictSlash(true)
	for i, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		endpoints[i] = route.Name + " : " + route.Pattern
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	services.WriteResponse(w, services.NewAPIResponse(endpoints))
}
