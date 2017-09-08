package main

import (
	"net/http"

	"github.com/jaguililla/goeasy/routing"
)

func main() {
	router := routing.Router()

	prefix := router.Subrouter("/prefix")
	prefix.Put(routing.Body("foo"))
	prefix.All(routing.Body("405 invalid method"))
	prefix.PathPut("/pput", hnd)

	router.PathGet("/other", routing.Body("Other"))
	router.All(routing.Body("999"))

	router.Serve(":8000")
}

func hnd(w http.ResponseWriter, r *http.Request) {
}
