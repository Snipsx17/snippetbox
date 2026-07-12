package main

import (
	"log"
	"net/http"
)

// handle home route "/"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from  Snippetbox app! 🚀"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	PORT := "4000"

	log.Print("Server running on localhost:" + PORT)
	err := http.ListenAndServe(":4000", mux)
	// Other way to run the server ↓↓
	// log.Fatal(http.ListenAndServe(":"+PORT, mux))
	log.Fatal(err)
}
