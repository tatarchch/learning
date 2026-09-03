package contractor

import (
	"context"
	"fmt"
)

type Service struct {
	repo repository
}

type CreateInput struct {
	Name string
	INN  string
}

func NewService(repo repository) *Service {
	if repo == nil {
		panic("nil repository")
	}

	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Contractor, error) {
	contractor, err := New(input.Name, input.INN)
	if err != nil {
		return Contractor{}, err
	}

	created, err := s.repo.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (Contractor, error) {
	if id <= 0 {
		return Contractor{}, ErrInvalidID
	}

	contractor, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Contractor{}, fmt.Errorf("find contractor by id: %w", err)
	}

	return contractor, nil
}

func (s *Service) Rename(ctx context.Context, id int64, name string) (Contractor, error) {
	if id <= 0 {
		return Contractor{}, ErrInvalidID
	}

	contractor, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Contractor{}, fmt.Errorf("find contractor by id: %w", err)
	}

	err = contractor.Rename(name)
	if err != nil {
		return Contractor{}, fmt.Errorf("rename contractor: %w", err)
	}

	contractor, err = s.repo.Update(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("update contractor: %w", err)
	}

	return contractor, nil
}
