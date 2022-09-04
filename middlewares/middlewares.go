package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// M1 is middleware 1
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.URL.RawQuery)
		next.ServeHTTP(w, r)
	})
}

// wrapper wraps http.Handler and returns httprouter.Handle
func Wrapper(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//pass httprouter.Params to request context
		ctx := context.WithValue(r.Context(), "params", ps)
		//call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
