package main

import (
	"fmt"
	"github.com/TV4/graceful"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/hostrouter"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

const DOMAIN = "versi.onl"

func BumpRouter(incType string) chi.Router {
	r := router()
	r.Handle("/{version}", BumpHandler(incType))
	return r
}

func router() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.NotFound(NotFoundHandler)

	r.HandleFunc("/health", func(_ http.ResponseWriter, _ *http.Request) {})
	return r
}

func main() {
	r2 := router()
	hr := hostrouter.New()
	for _, incType := range []string{"major", "minor", "patch"} {
		r2.Handle(fmt.Sprintf("/%s/{version}", incType), BumpHandler(incType))
		hr.Map(fmt.Sprintf("%s.%s", incType, DOMAIN), BumpRouter(incType))
	}
	hr.Map("*", r2)

	r := router()
	r.Mount("/", hr)
	graceful.ListenAndServe(&http.Server{Addr: ":8080", Handler: r})
}
