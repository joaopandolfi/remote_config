package vault

import (
	"encoding/json"
	"net/http"
	"oraculo/services"
	"oraculo/web/controllers"
	"oraculo/web/server"
	"strconv"

	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/utils"
)

type controller struct {
	s            *server.Server
	vaultService services.Vault
}

// New Vault controller
func New() controllers.Controller {
	return &controller{
		s:            nil,
		vaultService: services.NewVault(),
	}
}

func (c *controller) registerVault(w http.ResponseWriter, r *http.Request) {
	var vault Vault
	//err := handlers.SnakeCaseDecoder(r.Body).Decode(&vault)
	err := json.NewDecoder(r.Body).Decode(&vault)

	if err != nil {
		utils.Error("[RegisterVault] - unmarshaling", err.Error())
		handlers.RESTResponseError(w, map[string]string{"msg": "invalid body", "stack": err.Error()})
		return
	}

	privKey, err := c.vaultService.New(vault.ToModel())
	if err != nil {
		utils.CriticalError("[GetVault] - saving vault", err.Error())
		handlers.RESTResponseError(w, "problem to save vault")
		return
	}

	handlers.RESTResponse(w, map[string]string{"priv_key": privKey})
}

func (c *controller) getVault(w http.ResponseWriter, r *http.Request, public bool) {
	privKey := handlers.GetHeader(r, "key")
	q := handlers.GetQueryes(r)
	if privKey == "" {
		handlers.RESTResponseError(w, "invalid key")
		return
	}

	vault, err := c.vaultService.Get(privKey, public)
	if err != nil {
		utils.CriticalError("[GetVault] - recovering vault (public)", public, err.Error())
		handlers.RESTResponseError(w, "problem to recover vault")
		return
	}

	full, _ := strconv.ParseBool(q.Get("full"))
	if full {
		handlers.RESTResponse(w, vault)
		return
	}

	handlers.RESTResponse(w, vault.Metadata)
}

func (c *controller) getPrivate(w http.ResponseWriter, r *http.Request) {
	c.getVault(w, r, false)
}

func (c *controller) getPublic(w http.ResponseWriter, r *http.Request) {
	c.getVault(w, r, true)
}
