package router

import (
	"oraculo/config"
	"oraculo/web/controllers/auth"
	"oraculo/web/controllers/health"
	"oraculo/web/controllers/vault"
	"oraculo/web/server"

	"github.com/unrolled/secure"
)

// Router public struct
type Router struct {
	s *server.Server
}

// New Router
func New(s *server.Server) Router {
	return Router{s: s}
}

// Setup router
func (r *Router) Setup() {
	r.secure()

	health.New().SetupRouter(r.s)
	vault.New().SetupRouter(r.s)

	api := r.createSubRouter("/api")
	auth.New().SetupRouter(api)

}

// CreateSubRouter with path
func (r *Router) createSubRouter(path string) *server.Server {
	return &server.Server{
		R:      r.s.R.PathPrefix(path).Subrouter(),
		Config: r.s.Config,
	}
}

func (r *Router) secure() {
	secureMiddleware := secure.New(config.Get().Propertyes.Security.Options)
	r.s.R.Use(secureMiddleware.Handler)
}
