package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxCallback func(http.ResponseWriter, *http.Request)

// Call .
type Call struct {
	writer  http.ResponseWriter
	request *http.Request
}

// Response .
type Response struct {
	Code        int
	Body        string
	ContentType string
}

// Callback .
type Callback func(Call) (Response, error)

func (callback Callback) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if response, err := callback(Call{writer, request}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	} else {
		writer.WriteHeader(response.Code)
		writer.Header().Set("Content-Type", response.ContentType)
		fmt.Fprintln(writer, response.Body)
	}
}

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
func (router HTTPRouter) Get(block Callback) *mux.Route {
	return router.Handle("", block).Methods("GET")
}

// Put .
func (router HTTPRouter) Put(block Callback) *mux.Route {
	return router.Handle("", block).Methods("PUT")
}

// All .
func (router HTTPRouter) All(block Callback) *mux.Route {
	return router.Handle("", block)
}

// PathGet .
func (router HTTPRouter) PathGet(path string, block Callback) *mux.Route {
	return router.Handle(path, block).Methods("GET")
}

// PathPut .
func (router HTTPRouter) PathPut(path string, block Callback) *mux.Route {
	return router.Handle(path, block).Methods("PUT")
}

// PathAll .
func (router HTTPRouter) PathAll(path string, block Callback) *mux.Route {
	return router.Handle(path, block)
}

// Reply .
func Reply(message Response) Callback {
	return func(c Call) (Response, error) {
		return message, nil
	}
}

// Code .
func Code(code int) Callback {
	return Reply(Response{Code: code})
}

// Body .
func Body(message string) Callback {
	return Reply(Response{Code: 200, Body: message})
}

// CodeAndBody .
func CodeAndBody(code int, message string) Callback {
	return Reply(Response{Code: code, Body: message})
}
