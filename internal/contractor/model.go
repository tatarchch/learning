package contractor

import (
	"errors"
	"strings"
)

type Contractor struct {
	id   int
	name string
	inn  string
}

type CreateInput struct {
	name string
	inn  string
}

var (
	ErrNameRequired = errors.New("contractor name is required")
	ErrInnRequired  = errors.New("contractor inn is required")
)

// NewContractor validate name and inn strings
// and return contractor
func NewContractor(name, inn string) (Contractor, error) {
	name = normalizeName(name)
	inn = normalizeINN(inn)

	if name == "" {
		return Contractor{}, ErrNameRequired
	}

	if inn == "" {
		return Contractor{}, ErrInnRequired
	}

	return Contractor{name: name, inn: inn}, nil
}

func (c *Contractor) rename(name string) error {
	name = normalizeName(name)
	if name == "" {
		return ErrNameRequired
	}

	c.name = name
	return nil
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func normalizeINN(inn string) string {
	return strings.TrimSpace(inn)
}
