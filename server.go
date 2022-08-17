package main

import (
	"fmt"
	"log"
	"main/utilities"
	"net/http"
	"text/template"

	"github.com/BurntSushi/toml"
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
	PublicImportmaps   []importmap `toml:"public_importmaps"`
	PrivateImportmaps  []importmap `toml:"private_importmaps"`
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

func main() {
	// getting and parsing template files
	templateFiles := utilities.GetTemplates()
	tmpl, _ = template.ParseFiles(templateFiles...)

	// loading config
	loadConfig()

	router := httprouter.New()
	router.GET("/", Home)
	router.GET("/blog-home", BlogHome)
	router.GET("/posts/:link", BlogPost)
	router.GET("/posts", BlogPosts)
	router.ServeFiles("/static/*filepath", http.Dir("./public"))
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "home", map[string]interface{}{
		"Title": "Web app with Go std",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "bloghome", map[string]interface{}{
		"Title": "Blog -- Home",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	if err := tmpl.ExecuteTemplate(w, "blogpost", map[string]interface{}{
		"Title": "Blog -- Post Title",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BlogPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := tmpl.ExecuteTemplate(w, "blogposts", map[string]interface{}{
		"Title": "Blog -- Post Lists",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
