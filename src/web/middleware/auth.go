package middleware

import (
	"fmt"
	"net/http"

	"oraculo/config"
	"oraculo/services"

	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

// TokenHandler -
// @handler
// Intercept all transactions and check if is authenticated by token
func TokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Headers", "*")

			handlers.Response(w, "", 200)
			return
		}
		url := r.URL.String()

		token := handlers.GetHeader(r, "token")
		userID := handlers.GetHeader(r, "id")

		t, err := services.NewAuth().CheckToken(token)

		if !t.Authorized || err != nil || t.ID != userID {
			utils.Debug("[TokenHandler]", "Auth Error", url)
			handlers.Response(w, "Are you not my Preeecioouus", http.StatusForbidden)
			return
		}

		handlers.InjectHeader(r, "_xlevel", t.Permission)
		handlers.InjectHeader(r, "_xinstitution", t.Institution)
		handlers.InjectHeader(r, "_xid", t.ID)

		utils.Debug("[TokenHandler]", "Authenticated", url)
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - check if request have basic authentication
// @middleware
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := handlers.GetHeader(r, "Authentication")

		if token != fmt.Sprintf("Bearer %s", config.Get().BasicAuth) {
			handlers.RESTResponseWithStatus(w, "in valid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AuthTokenedProtection - Chain Logged handler to protect connections
// @middleware
// Uses session stored value `logged` to make a best gin of the world
// If is not connected, check token
func AuthTokenedProtection(f http.HandlerFunc) http.HandlerFunc {
	return handlers.Chain(f, TokenHandler)
}

// HandleToken -
func HandleToken(r *mux.Router, path string, f http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, AuthTokenedProtection(f)).Methods(methods...)
}

// HandleBasicAuth - handle basic authentication
func HandleBasicAuth(r *mux.Router, path string, f http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, BasicAuth(f)).Methods(methods...)
}

// Options allow to
func Options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}

	handlers.Response(w, "", 200)
}
