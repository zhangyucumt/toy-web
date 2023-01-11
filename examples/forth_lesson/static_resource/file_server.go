package main

import (
	"net/http"
	"path"
	"strings"
)

func main() {
	serve := http.FileServer(http.Dir("."))
	//http.Handle("/", serve)
	http.ListenAndServe(":8082", protectFiles(serve))
}

func protectFiles(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowed := ValidateDotFiles(r.URL.Path)
		if allowed {
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}

func ValidateDotFiles(fullPath string) bool {
	basePath := path.Base(fullPath)
	isPrivate := (strings.HasPrefix(basePath, "."))
	if isPrivate {
		return false
	} else {
		return true
	}
}
