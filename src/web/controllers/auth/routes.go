package auth

import (
	md "oraculo/web/middleware"
	"oraculo/web/server"
)

func (c *controller) SetupRouter(s *server.Server) {
	c.s = s

	c.s.R.Methods("OPTIONS").HandlerFunc(Options)
	c.s.R.HandleFunc("/token/check", c.checkToken).Methods("GET")
	md.HandleBasicAuth(s.R, "/token/new", c.newToken, "POST")
}
