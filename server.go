package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"main/build"
	"main/config"
	"main/handlers"
	"main/middlewares"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

var reloaded = false

func Live(ws *websocket.Conn) {
	// var received string
	for {
		if !reloaded {
			err := websocket.Message.Send(ws, "reload")
			log.Println("reload sent")
			if err != nil {
				panic(err)
			}
			reloaded = true
		}

		// break
	}
}

func serveEmbedded(h http.Handler) http.Handler {
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

func init() {
	// only do this in development environment
	if config.IsDevelopment() {
		build.BuildCSS()
	}
}

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// mux/router definition
	mux := http.NewServeMux()

	// public HTML route middleware stack
	publicHTMLStack := []middlewares.Middleware{
		middlewares.Logger,
	}

	// private JSON route middleware stack
	privateJSONStack := []middlewares.Middleware{
		middlewares.Logger,
		middlewares.VerifyToken,
	}

	// HTML routes
	mux.HandleFunc("/", middlewares.CompileMiddleware(handlers.Home, publicHTMLStack))
	mux.HandleFunc("/design", handlers.Design)
	mux.HandleFunc("/islands", handlers.Islands)

	// JSON REST posts API resources
	mux.HandleFunc("/api/posts", middlewares.CompileMiddleware(handlers.APIBlogPostsResource, privateJSONStack))

	// JSON REST tokens API resources
	mux.HandleFunc("/api/tokens", middlewares.CompileMiddleware(handlers.APITokensResources, publicHTMLStack))

	if config.IsDevelopment() {
		mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
		mux.Handle("/live", websocket.Handler(Live))
	} else {
		mux.Handle("/public/", serveEmbedded(http.FileServer(http.FS(embededPublic))))
	}

	log.Fatal(http.ListenAndServe(":8000", mux))
}
