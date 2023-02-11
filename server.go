package main

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log"
	"main/handlers"
	"net/http"

	"golang.org/x/net/websocket"
)

//go:embed templates
var embededTemplates embed.FS

//go:embed public
var embededPublic embed.FS

func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	http.HandleFunc("/", handlers.Home)

	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	public, err := fs.Sub(embededPublic, "public")
	if err != nil {
		panic(err)
	}
	http.Handle("/public", http.FileServer(http.FS(public)))
	http.Handle("/live", websocket.Handler(EchoServer))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// import (
// 	"embed"
// 	"html/template"
// 	"io"
// 	"io/fs"
// 	"log"
// 	"main/build"
// 	"main/config"
// 	"main/handlers"
// 	"main/middlewares"
// 	"net/http"

// 	// "github.com/gorilla/websocket"
// 	"github.com/julienschmidt/httprouter"
// 	"github.com/justinas/alice"
// 	"golang.org/x/net/websocket"
// )

// //go:embed templates
// var embededTemplates embed.FS

// //go:embed public
// var embededPublic embed.FS

// // var upgrader = websocket. // use default options

// func init() {
// 	build.BuildCSS()
// }

// func main() {
// 	// pre-parse templates, embedded in server binary
// 	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

// 	// router
// 	router := httprouter.New()

// 	// middlewares
// 	chain := alice.New(middlewares.Logger)
// 	// jwt validation middleware chain
// 	jwtChain := alice.New(middlewares.Logger, middlewares.VerifyToken)

// 	// HTML routes
// 	router.GET("/", middlewares.Wrapper(chain.ThenFunc(handlers.Home)))
// 	router.GET("/design", middlewares.Wrapper(chain.ThenFunc(handlers.Design)))
// 	router.GET("/islands", middlewares.Wrapper(chain.ThenFunc(handlers.Islands)))

// 	// JSON routes
// 	router.GET("/api/posts", middlewares.Wrapper(jwtChain.ThenFunc(handlers.APIBlogPosts)))
// 	router.POST("/api/posts", middlewares.Wrapper(jwtChain.ThenFunc(handlers.APICreateBlogPost)))
// 	// generate token
// 	router.POST("/api/tokens", middlewares.Wrapper(chain.ThenFunc(handlers.APICreateToken)))

// // static routes, embedded in server binary
// if public, err := fs.Sub(embededPublic, "public"); err == nil {
// 	if config.IsDevelopment() {
// 		// keep serving statiuc files from filesystem in development
// 		router.ServeFiles("/public/*filepath", http.Dir("public"))
// 	} else {
// 		fileServer := http.FileServer(http.FS(public))
// 		router.GET("/public/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 			w.Header().Set("Vary", "Accept-Encoding")
// 			w.Header().Set("Cache-Control", "public, max-age=7776000")
// 			r.URL.Path = p.ByName("filepath")
// 			fileServer.ServeHTTP(w, r)
// 		})
// 	}
// } else {
// 	panic(err)
// }

// 	// DEVELOPMENT only routes
// 	if config.IsDevelopment() {
// 		// router.GET("/live", live)
// 		router.Handle("/live", "", live)
// 	}

// 	// HTTP server
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

// func EchoServer(ws *websocket.Conn) {
// 	io.Copy(ws, ws)
// }

// func live(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// }

// func live(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}
// 	defer c.Close()

// 	for {
// 		mt, msg, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("received: %s", msg)
// 		err = c.WriteMessage(mt, []byte("reload"))
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// }
