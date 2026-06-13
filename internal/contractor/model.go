package contractor

import (
	"errors"
	"strings"
)

type Contractor struct {
	ID   int64
	Name string
	INN  string
}

type CreateInput struct {
	Name string
	INN  string
}

var (
	ErrNameRequired = errors.New("contractor name is required")
	ErrINNRequired  = errors.New("contractor inn is required")
)

// NewContractor validate name and inn strings
// and return Contractor
func NewContractor(name, inn string) (Contractor, error) {
	name = normalizeName(name)
	inn = normalizeINN(inn)

	if name == "" {
		return Contractor{}, ErrNameRequired
	}

	if inn == "" {
		return Contractor{}, ErrINNRequired
	}

	return Contractor{Name: name, INN: inn}, nil
}

func (c *Contractor) Rename(name string) error {
	name = normalizeName(name)
	if name == "" {
		return ErrNameRequired
	}

	c.Name = name
	return nil
}

func (c *Contractor) ChangeINN(inn string) error {
	inn = normalizeINN(inn)

	if inn == "" {
		return ErrINNRequired
	}

	return c.Rename(inn)
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func normalizeINN(inn string) string {
	return strings.TrimSpace(inn)
}
