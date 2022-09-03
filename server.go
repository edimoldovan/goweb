package main

import (
	"fmt"
	"log"
	"main/utilities"
	"net/http"
	"os"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type Post struct {
	Published string
	Title     string
	Url       string
	Image     string
	Lead      string
}

type importmap struct {
	Name string
	Path string
}

type tomlConfig struct {
	PostGresConnectURL string      `toml:"postgres_connect_url"`
	BaseUrl            string      `toml:"base_url"`
	BaseDomain         string      `toml:"base_domain"`
	Importmaps         []importmap `toml:"importmaps"`
}

var (
	config tomlConfig
)

func loadConfig() {
	f := "./config.toml"

	if _, err := toml.DecodeFile(f, &config); err != nil {
		log.Fatalln("Reading config failed", err)
	}

	// examples of config use
	// log.Println("PostGres URL:", config.PostGresConnectURL)

	// log.Println("Base URL:", config.BaseUrl)
	// log.Println("Jsimports:", config.Jsimports[0].Name)
	// log.Println("Jsimports:", config.Jsimports[0].Path)
}

var tmpl *template.Template
var upgrader = websocket.Upgrader{} // use default options

func main() {
	// read template files
	templateFiles := utilities.GetTemplates()
	// parse template files
	tmpl, _ = template.ParseFiles(templateFiles...)

	// config
	loadConfig()

	// route and routes
	router := httprouter.New()
	router.GET("/", Home)
	router.GET("/blog-home", BlogHome)
	router.GET("/posts/:link", BlogPost)
	router.GET("/posts", BlogPosts)
	router.ServeFiles("/static/*filepath", http.Dir("./public"))
	if os.Getenv("G_WEB_ENV") == "development" {
		router.GET("/live", live)
	}
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

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "home", map[string]interface{}{
		"Title":       "Web app with Go std",
		"Importmaps":  config.Importmaps,
		"Development": os.Getenv("G_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "bloghome", map[string]interface{}{
		"Title":       "Blog -- Home",
		"Importmaps":  config.Importmaps,
		"Development": os.Getenv("G_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	if err := tmpl.ExecuteTemplate(w, "blogpost", map[string]interface{}{
		"Title":       "Blog -- Post Title",
		"Importmaps":  config.Importmaps,
		"Development": os.Getenv("G_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "blogposts", map[string]interface{}{
		"Title":       "Blog -- Post Lists",
		"Importmaps":  config.Importmaps,
		"Development": os.Getenv("G_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
