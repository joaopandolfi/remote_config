package vault

import (
	md "oraculo/web/middleware"
	"oraculo/web/server"
)

func (c *controller) SetupRouter(s *server.Server) {
	c.s = s

	vault := s.R.PathPrefix("/vault").Subrouter()
	vault.HandleFunc("/public", c.getPublic).Methods("GET", "HEAD", "OPTIONS")
	vault.HandleFunc("/recover", c.getPrivate).Methods("GET")
	//md.HandleGandalfToken(vault, "/recover", c.getPrivate, "GET")
	md.HandleToken(vault, "/new", c.registerVault, "POST")

}
