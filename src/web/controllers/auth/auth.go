package auth

import (
	"net/http"

	"oraculo/web/controllers"
	"oraculo/web/server"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

type controller struct {
	s *server.Server
}

// New Auth Controller
func New() controllers.Controller {
	return &controller{
		s: nil,
	}
}

// CheckToken - validate JwtToken
func (c *controller) checkToken(w http.ResponseWriter, r *http.Request) {
	token := handlers.GetHeader(r, "token")
	q := handlers.GetQueryes(r)
	bearer := q.Get("bearer")

	if token != "" || bearer != "" {
		token = token + bearer
		t, err := utils.CheckJwtToken(token)
		if err == nil && t.Authorized {
			handlers.Response(w, map[string]interface{}{"logged": true, "data": t}, http.StatusOK)
			return
		}
	}

	handlers.Response(w, map[string]interface{}{"logged": false, "message": "Invalid Token"}, http.StatusOK)
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
