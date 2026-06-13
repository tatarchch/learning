package contractor

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var ErrInvalidId = errors.New("invalid ID. Id must be an integer and positive")

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

func (s *Service) Rename(ctx context.Context, id int64, name string) (Contractor, error) {
	if id <= 0 || reflect.TypeOf(id).Name() != "int64" {
		return Contractor{}, ErrInvalidId
	}

	contractor, err := s.repo.FindById(ctx, id)
	if err != nil {
		return Contractor{}, err
	}

	err = contractor.Rename(name)
	if err != nil {
		return Contractor{}, err
	}

	contractor, err = s.repo.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, err
	}

	return contractor, nil
}
