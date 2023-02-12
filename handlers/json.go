// curl -H 'Content-Type: application/json' -d '{ "email":"email","password":"password"}' -X POST http://localhost:8000/api/tokens
// curl -H 'Content-Type: application/json' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc0OTczMzAsImZvbyI6InNvbWV0aGluZyIsIm5iZiI6MTY0MDk5NTIwMH0.j8BkhToas63iamfzWSHP0HNdQCkK0nylLy8-XO-hzjY' -d '{ "published":"some date","title":"post title","url":"/some/url","image":"/some/image/url","lead":"some longer text"}' -X POST http://localhost:8000/api/posts

package handlers

import (
	"encoding/json"
	"net/http"
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

func APIBlogPostsResource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the resource.
		posts := []Post{
			{
				Published: "some date",
				Title:     "some title",
				Url:       "http://some.url",
			},
		}
		json.NewEncoder(w).Encode(posts)
	case http.MethodPost:
		// Create a new record.
		var post Post
		jsonErr := json.NewDecoder(r.Body).Decode(&post)
		if jsonErr != nil {
			http.Error(w, jsonErr.Error(), 400)
			return
		}
		json.NewEncoder(w).Encode(post)
	case http.MethodPut:
		// Update an existing record.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		// Remove the record.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func APITokensResources(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Serve the resource.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPost:
		// Create a new record.
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
	case http.MethodPut:
		// Update an existing record.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		// Remove the record.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
