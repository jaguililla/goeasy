package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	s := subrouter(r, "/prefix")
	put(s, body("foo"))
	all(s, body("405 invalid method"))

	http.ListenAndServe(":8000", r)
}

// Callback Route callback type
type Callback func(http.ResponseWriter, *http.Request)

func get(router *mux.Router, block Callback) {
	router.HandleFunc("", block).Methods("GET")
}

func put(router *mux.Router, block Callback) {
	router.HandleFunc("", block).Methods("PUT")
}

func all(router *mux.Router, block Callback) {
	router.HandleFunc("", block)
}

func subrouter(router *mux.Router, path string) *mux.Router {
	return router.PathPrefix(path).Subrouter()
}

func body(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	}
}
