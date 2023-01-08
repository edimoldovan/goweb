package main

import (
	"html/template"
	"log"
	"main/config"
	"main/handlers"
	"main/middlewares"
	"main/utilities"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// var templateFiles []string

// var Tmpl *template.Template
var upgrader = websocket.Upgrader{} // use default options

func main() {
	// config
	config.LoadConfig()
	// read template files
	templateFiles := utilities.GetTemplates()
	// parse template files
	handlers.Tmpl, _ = template.ParseFiles(templateFiles...)
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

	// static routes
	f := utilities.GetExecutable()
	if os.Getenv("G_WEB_ENV") == "development" {
		router.ServeFiles("/static/*filepath", http.Dir(f+"public"))
	} else {
		fileServer := http.FileServer(http.Dir(f + "public"))
		router.GET("/static/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			w.Header().Set("Vary", "Accept-Encoding")
			w.Header().Set("Cache-Control", "public, max-age=7776000")
			r.URL.Path = p.ByName("filepath")
			fileServer.ServeHTTP(w, r)
		})
	}

	// DEVELOPMENT only routes
	if os.Getenv("G_WEB_ENV") == "development" {
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
