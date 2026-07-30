package main

import (
	"encoding/json"
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
	w.Header().Add("Server", "GO")
	w.Write([]byte("Hello from  Snippetbox app! 🚀"))
}

func (h *SnippetHandler) View(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	resp := struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{
		Id:   id,
		Name: "Uberth",
	}

	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, `{"id": "%d", "name": "Uberth"}`, id)
	//w.Write([]byte(response))
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

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

func main() {
	mux := http.NewServeMux()

	// handler
	snippetHandler := NewSnippetHandler()

	//routes
	mux.HandleFunc("GET /{$}", snippetHandler.Home) // {$} restrict this route to strict matches on / only
	mux.HandleFunc("GET /snippet/view/{id}", snippetHandler.View)
	mux.HandleFunc("GET /snippet/create", snippetHandler.Create)
	mux.HandleFunc("POST /snippet/create", snippetHandler.CreatePost)

	PORT := "4000"

	log.Print("Server running on localhost:" + PORT)

	serverWrapper := LogRequest(mux)
	err := http.ListenAndServe(":4000", serverWrapper)
	// Other way to run the server ↓↓
	// log.Fatal(http.ListenAndServe(":"+PORT, mux))
	log.Fatal(err)
}
