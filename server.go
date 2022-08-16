package main

import (
	"fmt"
	"log"
	"main/utilities"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tmpl *template.Template

func main() {
	templateFiles := utilities.GetTemplates()
	tmpl, _ = template.ParseFiles(templateFiles...)
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
