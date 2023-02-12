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
	build.BuildCSS()
}

func main() {
	// pre-parse templates, embedded in server binary
	handlers.Tmpl = template.Must(template.ParseFS(embededTemplates, "templates/layouts/*.html", "templates/partials/*.html"))

	// router := http.NewServeMux()

	// middlewares
	// chain := alice.New(middlewares.Logger)s
	// // jwt validation middleware chain
	// jwtChain := alice.New(middlewares.Logger, middlewares.VerifyToken)

	http.Handle("/", middlewares.Logger(http.HandlerFunc(handlers.Home)))
	http.HandleFunc("/design", handlers.Design)
	http.HandleFunc("/islands", handlers.Islands)

	if config.IsDevelopment() {
		http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
		http.Handle("/live", websocket.Handler(Live))
	} else {
		http.Handle("/public/", serveEmbedded(http.FileServer(http.FS(embededPublic))))
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func main() {

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
