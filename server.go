package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"main/handlers"
	"main/middlewares"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// var Tmpl *template.Template
var upgrader = websocket.Upgrader{} // use default options

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// router
	router := httprouter.New()

	// middlewares
	chain := alice.New(middlewares.Logger)

	// HTML routes
	router.GET("/", middlewares.Wrapper(chain.ThenFunc(handlers.Home)))
	router.GET("/design", middlewares.Wrapper(chain.ThenFunc(handlers.Design)))
	router.GET("/islands", middlewares.Wrapper(chain.ThenFunc(handlers.Islands)))

	// JSON routes
	router.GET("/api/posts", middlewares.Wrapper(chain.ThenFunc(handlers.APIBlogPosts)))
	router.POST("/api/posts", middlewares.Wrapper(chain.ThenFunc(handlers.APICreateBlogPost)))
	router.POST("/api/tokens", middlewares.Wrapper(chain.ThenFunc(handlers.APICreateToken)))

	// static routes, embedded in server binary
	if public, err := fs.Sub(embededPublic, "public"); err == nil {
		if os.Getenv("GO_WEB_ENV") == "development" {
			router.Handler("GET", "/public/*filepath", http.StripPrefix("/public", http.FileServer(http.FS(public))))
		} else {
			fileServer := http.FileServer(http.FS(public))
			router.GET("/public/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				w.Header().Set("Vary", "Accept-Encoding")
				w.Header().Set("Cache-Control", "public, max-age=7776000")
				r.URL.Path = p.ByName("filepath")
				fileServer.ServeHTTP(w, r)
			})
		}
	} else {
		panic(err)
	}

	// DEVELOPMENT only routes
	if os.Getenv("GO_WEB_ENV") == "development" {
		router.GET("/live", live)
	}

	// HTTP server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func live(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("received: %s", msg)
		err = c.WriteMessage(mt, []byte("reload"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
