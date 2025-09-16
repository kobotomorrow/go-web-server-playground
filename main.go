package main

import (
	"go-web-server-playground/handler"
	"go-web-server-playground/resthandler"
	"go-web-server-playground/router"
	"net/http"
)

func main() {
	router := router.NewRouter()

	restUserHandler := &resthandler.RESTHandler{
		Index:  &handler.UserIndexHandler{},
		Create: &handler.UserCreateHandler{},
	}

	router.Handle("/users", restUserHandler)

	http.ListenAndServe(":8080", router)
}
