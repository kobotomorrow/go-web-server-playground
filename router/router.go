package router

import "net/http"

type Router struct {
	routes map[string]http.Handler
}

func (r *Router) Handle(path string, handler http.Handler) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.routes[req.URL.Path]
	if handler != nil {
		handler.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.Handler),
	}
}