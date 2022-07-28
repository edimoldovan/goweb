package router

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

func New() http.Handler {
	templateFiles := getTemplates()
	tmpl, _ = template.ParseFiles(templateFiles...)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexPage())
	return mux
}

func indexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "default", map[string]interface{}{
			"Title": "Web app with Go std",
		}); err != nil {
			fmt.Printf("ERR: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.Copy(w, &buf)
	}
}
