package middlewares

import (
	"devbook_app/src/cookies"
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s: %s - %s, from %s\n", time.Now().Format("02/01/2006 15:04:05"), r.Method, r.RequestURI, r.RemoteAddr)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.Read(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		fmt.Printf("%v\n", cookie)
		next(w, r)
	}
}
