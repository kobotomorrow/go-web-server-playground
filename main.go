package main

import (
	"fmt"
	"net/http"
)

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

type RESTHandler struct {
	Index  http.Handler
	Create http.Handler
}

func (rh *RESTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rh.Index.ServeHTTP(w, r)
	case "POST":
		rh.Create.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

type UserIndexHandler struct{}

func (h *UserIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Index")
}

type UserCreateHandler struct{}

func (h *UserCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Create")
}

func main() {
	router := NewRouter()

	restUserHandler := &RESTHandler{
		Index:  &UserIndexHandler{},
		Create: &UserCreateHandler{},
	}

	router.Handle("/users", restUserHandler)

	http.ListenAndServe(":8080", router)
}
