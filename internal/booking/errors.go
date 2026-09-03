package booking

import "errors"

var (
	ErrInvalidReference = errors.New("invalid booking reference")
	ErrInvalidAmount    = errors.New("invalid total amount")
	ErrNilBookedAt      = errors.New("bookedAt is null")
	//ErrNotFound         = errors.New("booking not found")
)
