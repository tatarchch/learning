package booking

import (
	"time"
)

type Booking struct {
	reference   string
	totalAmount int64
	bookedAt    time.Time
}

// New validates the reference, totalAmount and bookedAt and returns a Booking.
func New(reference string, totalAmount int64, bookedAt time.Time) (Booking, error) {

	if len(reference) != 6 {
		return Booking{}, ErrInvalidReference
	}

	if bookedAt.IsZero() {
		return Booking{}, ErrNilBookedAt
	}

	if totalAmount < 0 {
		return Booking{}, ErrInvalidAmount
	}

	return Booking{reference,
		totalAmount,
		bookedAt,
	}, nil
}

func (b Booking) Reference() string {
	return b.reference
}

func (b Booking) BookedAt() time.Time {
	return b.bookedAt
}

func (b Booking) TotalAmount() int64 {
	return b.totalAmount
}
