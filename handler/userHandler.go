package handler

import (
	"fmt"
	"net/http"
)

type UserIndexHandler struct{}

func (h *UserIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Index")
}

type UserCreateHandler struct{}

func (h *UserCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Create")
}