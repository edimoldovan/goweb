package handlers

import (
	"fmt"
	"html/template"
	"log"
	"main/config"
	"net/http"
	"os"
)

var Tmpl *template.Template

func DebugLog(value any) {
	log.Println("####################################################")
	log.Println(value)
}

func Home(w http.ResponseWriter, r *http.Request) {
	envConfig := config.EnvConfig()
	if err := Tmpl.ExecuteTemplate(w, "home", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  envConfig.Importmaps,
		"Development": os.Getenv("GO_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Design(w http.ResponseWriter, r *http.Request) {
	envConfig := config.EnvConfig()
	if err := Tmpl.ExecuteTemplate(w, "design", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  envConfig.Importmaps,
		"Development": os.Getenv("GO_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Islands(w http.ResponseWriter, r *http.Request) {
	envConfig := config.EnvConfig()
	if err := Tmpl.ExecuteTemplate(w, "islands", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  envConfig.Importmaps,
		"Development": os.Getenv("GO_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func VanillaMicroApps(w http.ResponseWriter, r *http.Request) {
	envConfig := config.EnvConfig()
	if err := Tmpl.ExecuteTemplate(w, "vanilla-microapps", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  envConfig.Importmaps,
		"Development": os.Getenv("GO_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Talk(w http.ResponseWriter, r *http.Request) {
	envConfig := config.EnvConfig()
	if err := Tmpl.ExecuteTemplate(w, "talk", map[string]interface{}{
		"Title":       "Visuals, interaction and shipping it all to the browsers",
		"Importmaps":  envConfig.Importmaps,
		"Development": os.Getenv("GO_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
