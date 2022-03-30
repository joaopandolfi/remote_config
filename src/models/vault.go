package models

import (
	"encoding/json"
	"oraculo/models/sanitizable"
	"oraculo/remotes/utils"
	"time"

	"golang.org/x/xerrors"
)

type Vault struct {
	Version     string
	Deleted     bool
	Token       string
	CreatorID   string
	System      string
	CreatedAt   time.Time
	Metadata    map[string]interface{} `bson:"-"`
	RawMetadata string                 `json:"-"`
	Public      bool
	sanitizable.Sanitilizable
}

func (m *Vault) cryptable() []*string {
	return []*string{
		&m.CreatorID, &m.RawMetadata,
	}
}

func (m *Vault) Crypt() error {
	b, err := json.Marshal(m.Metadata)
	if err != nil {
		return xerrors.Errorf("marshaling metadata: %w", err)
	}

	m.RawMetadata = utils.ToBase64(b)
	return m.Sanitize(m.cryptable())
}

func (m *Vault) Decrypt() error {
	err := m.Restore(m.cryptable())
	if err != nil {
		return xerrors.Errorf("decript vault: %w", err)
	}

	b, err := utils.FromBase64(m.RawMetadata)
	if err != nil {
		return xerrors.Errorf("recovering from base64: %w", err)
	}

	err = json.Unmarshal(b, &m.Metadata)
	if err != nil {
		return xerrors.Errorf("unmarshaling metadata: %w", err)
	}
	return nil
}
