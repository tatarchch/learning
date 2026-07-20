package contractor

import "errors"

var (
	ErrInvalidID        = errors.New("contractor ID must be positive")
	ErrContractNotFound = errors.New("contractor not found")
	ErrNameRequired     = errors.New("contractor name is required")
	ErrINNRequired      = errors.New("contractor inn is required")
)
