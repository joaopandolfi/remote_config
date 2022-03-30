package services

import (
	"encoding/json"
	"fmt"

	"oraculo/config"

	"github.com/joaopandolfi/blackwhale/remotes/request"
	"github.com/joaopandolfi/blackwhale/utils"
	"golang.org/x/xerrors"
)

// Auth interface system
type Auth interface {
	CheckToken(token string) (utils.Token, error)
}

type auth struct {
}

type receivedToken struct {
	Logged  bool        `json:"logged"`
	Message string      `json:"message"`
	Data    utils.Token `json:"data"`
}

// NewAuth service
func NewAuth() Auth {
	return &auth{}
}

// CheckToken in gandalf
func (s *auth) CheckToken(token string) (utils.Token, error) {
	var recToken receivedToken
	url := fmt.Sprintf("%s/rest/check/token", config.Get().File["GANDALF_URL"])
	b, err := request.GetWithHeader(url, map[string]string{"token": token})

	if err != nil {
		return utils.Token{Authorized: false}, xerrors.Errorf("checking JWT token: %w", err)
	}

	err = json.Unmarshal(b, &recToken)

	if err != nil {
		return utils.Token{Authorized: false}, xerrors.Errorf("unmarshaling response: %w", err)
	}

	return recToken.Data, err
}
