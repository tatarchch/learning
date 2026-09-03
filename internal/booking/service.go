package booking

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type repository interface {
	Save(ctx context.Context, contractor Booking) (Booking, error)
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
	if reference == "" {
		return Booking{}, ErrInvalidReference
	}

	booking, err := s.repo.FindByReference(ctx, reference)

	if errors.Is(err, sql.ErrNoRows) {
		return Booking{}, ErrNotFound
	}

	if err != nil {
		return Booking{}, fmt.Errorf("find booking: %w", err)
	}

	return booking, nil
}
