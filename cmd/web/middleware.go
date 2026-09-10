package main

import (
	"log"
	"net/http"
	"slices"
	"strings"
)

type Middleware func(http.Handler) http.Handler

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func DisableFileServerListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasSuffix(path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
* Accept any number of parameters of type Middleware and expand it,
* it should be and array or slice
 */
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	/*
	* slices.Backward change de order of the slice
	* loop in all the slice create by the parameter middlewares that is a variadic param
	 */
	for _, middleware := range slices.Backward(middlewares) {
		// override the param h and chain the middlewares in the same order received
		h = middleware(h)
	}
	return h
}

var Middlewares = []Middleware{LogRequest, DisableFileServerListing}
