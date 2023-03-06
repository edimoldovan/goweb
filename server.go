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

// private JSON route middleware stack
var privateJSONStack = []middlewares.Middleware{
	middlewares.Logger,
	middlewares.VerifyToken,
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

	router.Routes = []router.Route{
		// HTML routes
		router.CreateRoute("GET", "/", middlewares.CompileMiddleware(handlers.Home, publicHTMLStack)),
		router.CreateRoute("GET", "/design", handlers.Design),
		router.CreateRoute("GET", "/islands", handlers.Islands),
		router.CreateRoute("GET", "/vanilla-microapps", handlers.VanillaMicroApps),
		// JSON API routes
		router.CreateRoute("GET", "/api/posts", middlewares.CompileMiddleware(handlers.APIBlogPostsResource, privateJSONStack)),
	}

	// only do this in development environment
	if config.IsDevelopment() {
		build.BuildCSS()

		router.Routes = append(router.Routes, router.CreateRoute("GET", "/live", http.HandlerFunc(websocket.Handler(Live).ServeHTTP)))
		staticServer = http.FileServer(http.Dir("./public"))
	} else {
		staticServer = http.FileServer(http.FS(embededPublic))
	}
	router.Routes = append(router.Routes, router.CreateRoute("GET", "/public/.*", http.StripPrefix("/public/", staticServer).ServeHTTP))
}

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// mux/router definition
	mux := http.HandlerFunc(router.Serve)

	// start the server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
