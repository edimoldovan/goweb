package main

import (
	"embed"
	"html/template"
	"log"
	"main/build"
	"main/config"
	"main/handlers"
	"main/middlewares"
	"main/router"
	"net/http"

	"golang.org/x/net/websocket"
)

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

var reloaded = false

// public HTML route middleware stack
var publicHTMLStack = []middlewares.Middleware{
	middlewares.Logger,
}

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

func init() {
	var staticServer http.Handler
	var stripPrefix string

	router.Routes = []router.Route{
		// HTML routes
		router.CreateRoute("GET", "/", middlewares.CompileMiddleware(handlers.Home, publicHTMLStack)),
		router.CreateRoute("GET", "/design", middlewares.CompileMiddleware(handlers.Design, publicHTMLStack)),
		router.CreateRoute("GET", "/islands", middlewares.CompileMiddleware(handlers.Islands, publicHTMLStack)),
		router.CreateRoute("GET", "/vanilla-microapps", middlewares.CompileMiddleware(handlers.VanillaMicroApps, publicHTMLStack)),
		router.CreateRoute("GET", "/talk", middlewares.CompileMiddleware(handlers.Talk, publicHTMLStack)),
	}

	// only do this in development environment
	if config.IsDevelopment() {
		build.BuildCSS()
		router.Routes = append(router.Routes, router.CreateRoute("GET", "/live", http.HandlerFunc(websocket.Handler(Live).ServeHTTP)))
		staticServer = http.FileServer(http.Dir("./public"))
		stripPrefix = "/public/"
	} else {
		staticServer = http.FileServer(http.FS(embededPublic))
		stripPrefix = "/"
	}
	router.Routes = append(router.Routes, router.CreateRoute("GET", "/public/.*", http.StripPrefix(stripPrefix, staticServer).ServeHTTP))
}

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// mux/router definition
	mux := http.HandlerFunc(router.Serve)

	// start the server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
