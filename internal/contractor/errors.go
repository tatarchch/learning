package contractor

import "errors"

var (
	ErrInvalidID    = errors.New("contractor ID must be positive")
	ErrNameRequired = errors.New("contractor name is required")
	ErrINNRequired  = errors.New("contractor inn is required")
)
