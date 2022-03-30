package controllers

import (
	"oraculo/web/server"
)

// Controller public contract
type Controller interface {
	SetupRouter(s *server.Server)
}
