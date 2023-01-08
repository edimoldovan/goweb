package handlers

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"main/config"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/go-session/session"
)

var Tmpl *template.Template

func DebugLog(value any) {
	log.Println("####################################################")
	log.Println(value)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "home", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  config.Config.Importmaps,
		"Development": os.Getenv("G_WEB_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(context.Background(), w, r)
	if err != nil {
		log.Println(err)
		return
	}
	store.Flush()
	err = store.Save()
	if err != nil {
		fmt.Fprint(w, err)
		// return
	}
	rndstring := strconv.Itoa(rand.Intn(1000000))
	http.Redirect(w, r, fmt.Sprintf(`/?rnd=%s`, rndstring), http.StatusFound)
}

func Design(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "design", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  config.Config.Importmaps,
		"Development": os.Getenv("DSGN_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Islands(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.ExecuteTemplate(w, "islands", map[string]interface{}{
		"Title":       "Web app with Go",
		"Importmaps":  config.Config.Importmaps,
		"Development": os.Getenv("DSGN_ENV") == "development",
	}); err != nil {
		fmt.Printf("ERR: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
