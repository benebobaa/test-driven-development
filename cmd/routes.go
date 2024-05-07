package main

import (
	"net/http"
	"test-driven-development/pkg/domain"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(service *domain.Service) *Router {
	router := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	router.addRoute("POST", "/users", service.CreateUser)

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, methodExists := handlers[req.Method]; methodExists {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *Router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}
