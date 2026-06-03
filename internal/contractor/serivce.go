package contractor

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNameRequired = errors.New("contractor name is required")
	ErrInnRequired  = errors.New("contractor inn is required")
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
	trimName := strings.TrimSpace(input.Name)
	trimInn := strings.TrimSpace(input.INN)

	if trimName == "" {
		return Contractor{}, ErrNameRequired
	}

	if trimInn == "" {
		return Contractor{}, ErrInnRequired
	}

	contractor := Contractor{
		Name: trimName,
		INN:  trimInn,
	}

	created, err := s.r.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}
