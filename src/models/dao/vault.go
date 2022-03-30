package dao

import (
	"oraculo/models"
	"oraculo/remotes/mongo"

	"golang.org/x/xerrors"
	"gopkg.in/mgo.v2/bson"
)

// Vault contract
type Vault interface {
	New(v models.Vault) error
	Get(token string, public bool) (*models.Vault, error)
	Update(token string, v models.Vault) error
}

type vault struct {
	mongo.Dao
}

// NewVault constructor
func NewVault() Vault {
	return &vault{
		Dao: mongo.Dao{Collection: "vault"},
	}
}

func (d *vault) New(v models.Vault) error {
	return d.Dao.Save(v)
}

func (d *vault) Get(token string, public bool) (*models.Vault, error) {
	var result models.Vault
	col, err := d.GetCollection()
	if err != nil {
		return nil, xerrors.Errorf("getting mongo connection %w", err)
	}

	err = col.Find(bson.M{"token": token, "public": public}).One(&result)
	if err != nil {
		return nil, xerrors.Errorf("Get vault: %w", err)
	}

	return &result, nil
}

func (d *vault) Update(token string, v models.Vault) error {
	col, err := d.GetCollection()
	if err != nil {
		return xerrors.Errorf("getting mongo connection %w", err)
	}

	err = col.Update(bson.M{"token": token}, v)
	if err != nil {
		return xerrors.Errorf("updating vault: %w", err)
	}
	return nil
}
