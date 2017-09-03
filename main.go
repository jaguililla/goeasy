package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Call .
type Call struct {
	writer  http.ResponseWriter
	request *http.Request
}

// Response .
type Response struct {
	code        int
	body        string
	contentType string
}

// Callback .
type Callback func(http.ResponseWriter, *http.Request)

// Callback2 .
type Callback2 func(Call) (Response, error)

func (cb Callback2) ServeHTTP(http.ResponseWriter, *http.Request) {}

// HTTPRouter .
type HTTPRouter struct {
	*mux.Router
}

// Router .
func Router() HTTPRouter {
	return HTTPRouter{mux.NewRouter()}
}

// Serve .
func (router HTTPRouter) Serve(address string) {
	http.ListenAndServe(address, router)
}

// Subrouter .
func (router HTTPRouter) Subrouter(path string) HTTPRouter {
	return HTTPRouter{router.PathPrefix(path).Subrouter()}
}

// Get .
func (router HTTPRouter) Get(block Callback) {
	router.HandleFunc("", block).Methods("GET")
}

// Put .
func (router HTTPRouter) Put(block Callback) {
	router.HandleFunc("", block).Methods("PUT")
}

// All .
func (router HTTPRouter) All(block Callback) {
	router.HandleFunc("", block)
}

// PathGet .
func (router HTTPRouter) PathGet(path string, block Callback) {
	router.HandleFunc(path, block).Methods("GET")
}

// PathPut .
func (router HTTPRouter) PathPut(path string, block Callback) {
	router.HandleFunc(path, block).Methods("PUT")
}

// PathAll .
func (router HTTPRouter) PathAll(path string, block Callback) {
	router.HandleFunc(path, block)
}

// Body .
func Body(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
	}
}

func f() error {
	return errors.New("fatal")
}

func f2() error {
	if err := f(); err != nil {
		return errors.New("Fail again")
	}

	return nil
}

func main() {
	router := Router()

	prefix := router.Subrouter("/prefix")
	//router.NotFoundHandler = http.Handler()
	prefix.Put(Body("foo"))
	//prefix.All(Body("405 invalid method"))
	prefix.PathPut("/pput", func(w http.ResponseWriter, r *http.Request) {
	})

	router.PathGet("/other", Body("Other"))
	router.All(Body("999"))

	router.Serve(":8000")
}
