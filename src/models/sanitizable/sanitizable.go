package sanitizable

import (
	"oraculo/config"
	"oraculo/remotes/utils/aes"

	"golang.org/x/xerrors"
)

// Sanitilizable - public struct to implement sanitization by criptography
type Sanitilizable struct {
	Sanitized bool `json:"sanitized" `
}

// Crypt received value
func Crypt(val string) (string, error) {
	encVal, err := aes.Encrypt(config.Get().Server.Security.AESKey, val)
	if err != nil {
		return "", xerrors.Errorf("encrypting: %w", err)
	}
	return encVal, nil
}

func Decrypt(val string) (string, error) {
	encVal, err := aes.Decrypt(config.Get().Server.Security.AESKey, val)
	if err != nil {
		return "", xerrors.Errorf("restoring: %v", err)
	}

	return encVal, nil
}

func (m *Sanitilizable) Sanitize(vals []*string) error {
	if m.Sanitized {
		return nil
	}

	for i, val := range vals {
		encVal, err := aes.Encrypt(config.Get().Server.Security.AESKey, *val)
		if err != nil {
			return xerrors.Errorf("encrypting %d: %v", i, err)
		}
		*vals[i] = encVal
	}
	m.Sanitized = true
	return nil
}

func (m *Sanitilizable) Restore(vals []*string) error {
	if !m.Sanitized {
		return nil
	}

	for i, val := range vals {
		encVal, err := aes.Decrypt(config.Get().Server.Security.AESKey, *val)
		if err != nil {
			return xerrors.Errorf("restoring %v", err)
		}
		*vals[i] = encVal
	}
	m.Sanitized = false
	return nil
}
