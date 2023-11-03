package main

import (
	"log"
	"net/http"

	"github.com/shortboard/go-htmx-templ-demo/assets"
	"github.com/shortboard/go-htmx-templ-demo/components"
	"github.com/shortboard/go-htmx-templ-demo/services"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	counter := services.NewCountService()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(components.Home(counter.GetCount())).ServeHTTP(w, r)
	})
	r.Get("/api/click", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(components.Counter(counter.GetCount())).ServeHTTP(w, r)
	})
	assets.Mount(r)
	log.Fatal(http.ListenAndServe(":4000", r))
}
