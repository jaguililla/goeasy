package main

import (
	r "github.com/jaguililla/goeasy/routing"
)

func main() {
	router := r.Router()

	router.PathGet("/other", r.Body("Other"))
	router.PathGet("/root", r.Body("Rooted"))

	prefix := router.Subrouter("/prefix")

	prefix.Put(r.Body("foo"))
	prefix.PathPut("/put", r.CodeAndBody(200, "response"))
	prefix.PathGet("/matched", func(r.Call) (r.Response, error) {
		return r.Response{Code: 200, Body: "Matched!"}, nil
	})

	prefix.All(r.CodeAndBody(405, "Invalid method"))

	router.Serve(":8000")
}
