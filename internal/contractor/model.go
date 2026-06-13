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
	var c Contractor

	if err := c.Rename(name); err != nil {
		return Contractor{}, err
	}

	if err := c.ChangeINN(inn); err != nil {
		return Contractor{}, err
	}

	return c, nil
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

	c.INN = inn
	return nil
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func normalizeINN(inn string) string {
	return strings.TrimSpace(inn)
}
