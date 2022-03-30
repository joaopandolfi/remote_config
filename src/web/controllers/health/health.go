package health

import (
	"net/http"

	"oraculo/web/controllers"
	"oraculo/web/server"

	"github.com/joaopandolfi/blackwhale/handlers"
)

// --- Health ---

type controller struct {
	s *server.Server
}

// New Health controller
func New() controllers.Controller {
	return &controller{
		s: nil,
	}
}

// Health route
func (c *controller) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	handlers.Response(w, true, http.StatusOK)
}
