package booking

import (
	"context"
	"fmt"
	"time"
)

type repository interface {
	Save(ctx context.Context, booking Booking) (Booking, error)
	FindByReference(ctx context.Context, reference string) (Booking, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	if repo == nil {
		panic("nil repository")
	}

	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context,
	reference string,
	totalAmount int,
	bookedAt time.Time,
) (Booking, error) {
	booking, err := New(reference, totalAmount, bookedAt)
	if err != nil {
		return Booking{}, err
	}

	created, err := s.repo.Save(ctx, booking)
	if err != nil {
		return Booking{}, fmt.Errorf("save booking: %w", err)
	}

	return created, nil
}

func (s *Service) GetByReference(ctx context.Context, reference string) (Booking, error) {
	if len(reference) != 6 {
		return Booking{}, ErrInvalidReference
	}

	booking, err := s.repo.FindByReference(ctx, reference)

	if err != nil {
		return Booking{}, fmt.Errorf("find booking: %w", err)
	}

	return booking, nil
}
