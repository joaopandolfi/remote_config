package services

import (
	"oraculo/models"
	"oraculo/models/dao"
	"oraculo/models/sanitizable"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type Vault interface {
	New(v models.Vault) (string, error)
	Get(token string, public bool) (*models.Vault, error)
	Update(token string, v models.Vault) error
}

type vault struct {
	vaultDao dao.Vault
}

func NewVault() Vault {
	return &vault{
		vaultDao: dao.NewVault(),
	}
}

func (s *vault) New(v models.Vault) (string, error) {
	v.CreatedAt = time.Now()
	token := uuid.New()

	tk, err := sanitizable.Crypt(token.String())
	if err != nil {
		return "", xerrors.Errorf("generating token: %w", err)
	}

	v.Token = tk

	err = v.Crypt()
	if err != nil {
		return "", xerrors.Errorf("crypting vault: %w", err)
	}

	err = s.vaultDao.New(v)
	if err != nil {
		return "", xerrors.Errorf("saving vault on database: %w", err)
	}

	return tk, nil
}

func (s *vault) Get(token string, public bool) (*models.Vault, error) {
	v, err := s.vaultDao.Get(token, public)
	if err != nil {
		return nil, xerrors.Errorf("getting vault: %w", err)
	}
	err = v.Decrypt()
	if err != nil {
		return nil, xerrors.Errorf("decrypt vault: %w", err)
	}
	return v, nil
}

func (s *vault) Update(token string, v models.Vault) error {
	v.CreatedAt = time.Now()
	err := v.Crypt()
	if err != nil {
		return xerrors.Errorf("crypting vault: %w", err)
	}

	return s.vaultDao.Update(token, v)
}
