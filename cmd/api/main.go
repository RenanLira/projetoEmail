package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {

		name := chi.URLParam(r, "name")

		w.Write([]byte("Hello World " + name))
	})

	http.ListenAndServe(":8080", r)
}
