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
	w.Write([]byte("Hello from  Snippetbox app! 🚀"))
}

func (h *SnippetHandler) View(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	response := fmt.Sprintf("View snippet with ID: %d", id)
	w.Write([]byte(response))
}

func (h *SnippetHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet"))
}

func (h *SnippetHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Create snippet POST"))
}

// middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

var snippetHandler = NewSnippetHandler()

func main() {
	mux := http.NewServeMux()

	//routes
	mux.HandleFunc("GET /{$}", snippetHandler.Home) // {$} restrict this route to strict matches on / only
	mux.HandleFunc("GET /snippet/view/{id}", snippetHandler.View)
	mux.HandleFunc("GET /snippet/create", snippetHandler.Create)
	mux.HandleFunc("POST /snippet/create", snippetHandler.CreatePost)

	PORT := "4000"

	log.Print("Server running on localhost:" + PORT)
	err := http.ListenAndServe(":4000", LogRequest(mux))
	// Other way to run the server ↓↓
	// log.Fatal(http.ListenAndServe(":"+PORT, mux))
	log.Fatal(err)
}
