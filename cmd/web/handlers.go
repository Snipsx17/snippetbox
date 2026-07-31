package main

import (
	"encoding/json"
	"html/template"
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

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
		"./ui/html/partials/footer.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

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
