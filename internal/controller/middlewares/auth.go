package middlewares

import (
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Not Allowed.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "pismo" && password == "12345678"
}
