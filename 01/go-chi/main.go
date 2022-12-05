package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World Response"))
}

func main() {
	mux := chi.NewRouter()
	mux.Get("/hello", helloHandler)

	http.ListenAndServe(":8080", mux)
}
