package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/utilities"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type Post struct {
	Published string `json:"published"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Image     string `json:"image"`
	Lead      string `json:"lead"`
}

type JWTToken struct {
	Token string `json:"token"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

var hmacSampleSecret = []byte("someSecret") // TODO: put this key in safe place and use proper key

func loadConfig() {
	f := "./config.toml"

	if _, err := toml.DecodeFile(f, &config); err != nil {
		log.Fatalln("Reading config failed", err)
	}

	// examples of config use
	// log.Println("PostGres URL:", config.PostGresConnectURL)
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

	// router
	router := httprouter.New()

	// HTML routes
	router.GET("/", Home)
	router.GET("/blog-home", BlogHome)
	router.GET("/posts/:link", BlogPost)
	router.GET("/posts", BlogPosts)

	// JSON routes
	router.GET("/api/posts", APIBlogPosts)
	router.POST("/api/posts", APICreateBlogPost)
	router.POST("/api/tokens", APICreateToken)

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

func APIBlogPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	posts := []Post{
		{
			Published: "some date",
			Title:     "some title",
			Url:       "http://some.url",
		},
	}
	json.NewEncoder(w).Encode(posts)
}

func APICreateBlogPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")

	token, tokenErr := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(tokenErr)
	}

	var post Post
	jsonErr := json.NewDecoder(r.Body).Decode(&post)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), 400)
		return
	}
	log.Println(post)
	json.NewEncoder(w).Encode(post)
}

func APICreateToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user User
	var token JWTToken
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if user.Email == "email" && user.Password == "password" {
		generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC).Unix(),
			"exp": time.Now().Add(time.Minute * 30).Unix(),
		})
		token.Token, err = generatedToken.SignedString(hmacSampleSecret)
	} else {
		token.Token = "wrong login credentials"
	}
	json.NewEncoder(w).Encode(token)
}
