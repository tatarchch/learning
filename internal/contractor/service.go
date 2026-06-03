package contractor

import (
	"context"
	"fmt"
)

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	if repo == nil {
		panic("nil repository")
	}

	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Contractor, error) {
	contractor, err := NewContractor(input.Name, input.INN)
	if err != nil {
		return Contractor{}, err
	}

	created, err := s.repo.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}
