package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type SnippetHandler struct {
}

func NewSnippetHandler() *SnippetHandler {
	return &SnippetHandler{}
}

// handle home route "/"
func (h *SnippetHandler) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)
	w.Write([]byte("Hello from  Snippetbox app! 🚀"))
}

func (h *SnippetHandler) ViewSnippet(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	response := fmt.Sprintf("View snippet with ID: %d", id)
	w.Write([]byte(response))
}

func (h *SnippetHandler) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)
	w.Write([]byte("Create snippet"))
}

var Handler = NewSnippetHandler()

func main() {
	mux := http.NewServeMux()

	//routes
	mux.HandleFunc("/{$}", Handler.Home) // {$} restrict this route to strict matches on / only
	mux.HandleFunc("/snippet/view/{id}", Handler.ViewSnippet)
	mux.HandleFunc("/snippet/create", Handler.CreateSnippet)

	PORT := "4000"

	log.Print("Server running on localhost:" + PORT)
	err := http.ListenAndServe(":4000", mux)
	// Other way to run the server ↓↓
	// log.Fatal(http.ListenAndServe(":"+PORT, mux))
	log.Fatal(err)
}
