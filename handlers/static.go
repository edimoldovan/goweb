package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ServeEmbedded(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		if r.URL.Path == "/" {
			r.URL.Path = fmt.Sprintf("/%s/", "public")
			log.Println(r.URL.Path)
		} else {
			b := strings.Split(r.URL.Path, "/")[1]
			if b != "public" {
				r.URL.Path = fmt.Sprintf("/%s%s", "public", r.URL.Path)
			}
		}
		h.ServeHTTP(w, r)
	})
}
