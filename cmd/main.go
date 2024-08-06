package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"turlarion.ru/url-shortener/internal/api"
)

func main() {
	r := chi.NewRouter()
	r.Get("/save", api.Save)
	r.Get("/hw", func(w http.ResponseWriter, r *http.Request) {
		println("HELLO WORLD!")
	})
	http.ListenAndServe(":8080", r)
}
