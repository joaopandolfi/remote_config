package auth

import (
	"oraculo/web/server"
)

func (c *controller) SetupRouter(s *server.Server) {
	c.s = s

	c.s.R.Methods("OPTIONS").HandlerFunc(Options)
	c.s.R.HandleFunc("/check/token", c.checkToken).Methods("GET")
}
