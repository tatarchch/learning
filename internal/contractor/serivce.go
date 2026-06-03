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

func (s *Service) create(ctx context.Context, input createInput) (Contractor, error) {
	contractor, err := newContractor(input.name, input.inn)
	if err != nil {
		return Contractor{}, err
	}

	created, err := s.r.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}
