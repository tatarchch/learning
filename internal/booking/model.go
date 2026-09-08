package booking

import (
	"maps"
	"slices"
	"time"
)

type Booking struct {
	reference   string
	totalAmount int64
	bookedAt    time.Time
	metadata    map[string]string
	passengers  []string
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
		make(map[string]string),
		make([]string, 0, 4),
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

func (b *Booking) SetMetadata(key, value string) {
	b.metadata[key] = value
}

func (b Booking) Metadata(key string) string {
	return b.metadata[key]
}

func (b Booking) Clone() Booking {
	metadata := make(map[string]string)
	maps.Copy(metadata, b.metadata)
	passenger := slices.Clone(b.passengers)

	return Booking{
		b.reference,
		b.totalAmount,
		b.bookedAt,
		metadata,
		passenger,
	}
}

func (b *Booking) ClearMetadata() {
	b.metadata = make(map[string]string)
}

func (b Booking) ClearedMetadata() Booking {
	b.metadata = make(map[string]string)
	return b
}

func (b *Booking) AddPassenger(name string) {
	b.passengers = append(b.passengers, name)
}

func (b Booking) Passengers() []string {
	return b.passengers
}
