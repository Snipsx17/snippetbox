package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// handler
	snippetHandler := NewSnippetHandler()

	// file server
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	//routes
	/*
	* {$} restrict this route to strict matches on / only
	 */
	mux.HandleFunc("GET /{$}", snippetHandler.Home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetHandler.View)
	mux.HandleFunc("GET /snippet/create", snippetHandler.Create)
	mux.HandleFunc("POST /snippet/create", snippetHandler.CreatePost)

	PORT := "4000"

	log.Print("Server running on localhost:" + PORT)

	// middlewares wrapper
	serverWrapper := Chain(mux, Middlewares...)

	// run server
	/*
	* Other way to run the server ↓↓
	* log.Fatal(http.ListenAndServe(":"+PORT, mux))
	 */
	err := http.ListenAndServe(":4000", serverWrapper)
	log.Fatal(err)
}
