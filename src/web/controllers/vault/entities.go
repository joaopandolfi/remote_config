package vault

import "oraculo/models"

type Vault struct {
	Version   string                 `json:"version"`
	Metadata  map[string]interface{} `json:"metadata"`
	System    string                 `json:"system"`
	CreatorID string                 `json:"creator_id"`
	Public    bool                   `json:"public"`
}

func (v *Vault) ToModel() models.Vault {
	return models.Vault{
		Metadata:  v.Metadata,
		System:    v.System,
		CreatorID: v.CreatorID,
		Version:   v.Version,
		Public:    v.Public,
	}
}
