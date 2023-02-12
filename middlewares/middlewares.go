package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("someSecret") // TODO: put this key in safe place and use proper secret

// type for chaining
type Middleware func(http.HandlerFunc) http.HandlerFunc

// basically thisd is middleware chaining
func CompileMiddleware(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			tokenArray := strings.Split(strings.Trim(r.Header.Get("Authorization"), " "), " ")
			parsedToken, tokenErr := jwt.Parse(tokenArray[1], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return hmacSampleSecret, nil
			})

			if tokenErr != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
				// store the claims in the request as a context obj
				r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
			} else {
				fmt.Println(tokenErr)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
