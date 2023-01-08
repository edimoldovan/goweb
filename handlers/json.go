package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("someSecret") // TODO: put this key in safe place and use proper secret

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

func APIBlogPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{
			Published: "some date",
			Title:     "some title",
			Url:       "http://some.url",
		},
	}
	json.NewEncoder(w).Encode(posts)
}

func APICreateBlogPost(w http.ResponseWriter, r *http.Request) {
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

func APICreateToken(w http.ResponseWriter, r *http.Request) {
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
