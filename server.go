package main

import (
	"log"
	"main/config"
	"main/handlers"
	"main/utilities"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

// var templateFiles []string

// var Tmpl *template.Template
var upgrader = websocket.Upgrader{} // use default options

func main() {
	// read template files
	templateFiles := utilities.GetTemplates()
	// parse template files
	handlers.Tmpl, _ = template.ParseFiles(templateFiles...)

	// config
	config.LoadConfig()

	// router
	router := httprouter.New()

	// HTML routes
	router.GET("/", handlers.Home)
	router.GET("/blog-home", handlers.BlogHome)
	router.GET("/posts/:link", handlers.BlogPost)
	router.GET("/posts", handlers.BlogPosts)

	// JSON routes
	router.GET("/api/posts", handlers.APIBlogPosts)
	router.POST("/api/posts", handlers.APICreateBlogPost)
	router.POST("/api/tokens", handlers.APICreateToken)

	// static routes
	router.ServeFiles("/static/*filepath", http.Dir("./public"))

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
