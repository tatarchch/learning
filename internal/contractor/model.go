package contractor

import (
	"errors"
	"strings"
)

type Contractor struct {
	id          int64
	name        string
	inn         string
	description string
}

var (
	ErrNameRequired = errors.New("contractor name is required")
	ErrINNRequired  = errors.New("contractor inn is required")
)

// NewContractor validates the name and INN and returns a Contractor.
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

func (c Contractor) ID() int64 {
	return c.id
}

func (c Contractor) Name() string {
	return c.name
}

func (c Contractor) INN() string {
	return c.inn
}

func (c Contractor) Description() string {
	return c.description
}

func (c *Contractor) Rename(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrNameRequired
	}

	c.name = name
	return nil
}

func (c *Contractor) ChangeINN(inn string) error {
	inn = strings.TrimSpace(inn)

	if inn == "" {
		return ErrINNRequired
	}

	c.inn = inn
	return nil
}

func (c *Contractor) ChangeDescription(description string) {
	description = strings.TrimSpace(description)

	c.description = description
	return
}
