package contractor

import (
	"context"
	"fmt"
)

type Service struct {
	r repository
}

func NewService(r repository) *Service {
	if r == nil {
		panic("nil repository")
	}

	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Contractor, error) {
	contractor, err := NewContractor(input.name, input.inn)
	if err != nil {
		return Contractor{}, err
	}

	created, err := s.r.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}
