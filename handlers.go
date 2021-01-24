package main

import (
	"github.com/Masterminds/semver"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func BumpHandler(incType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		version, err := semver.NewVersion(chi.URLParam(r, "version"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		v := version.IncMinor() // default is minor
		if incType == "major" {
			v = version.IncMajor()
		} else if incType == "patch" {
			v = version.IncPatch()
		}

		_, _ = w.Write([]byte(v.String()))
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		panic(err)
	}
	if err := tmpl.ExecuteTemplate(w, "404.html", DOMAIN); err != nil {
		panic(err)
	}
}
