package contractor

import (
	"context"
	"errors"
	"fmt"
)

var ErrInvalidID = errors.New("contractor ID must be positive")

type Service struct {
	creatorRepo creatorRepository
	renamerRepo renamerRepository
}

type CreateInput struct {
	Name string
	INN  string
}

func NewService(creatorRepo creatorRepository, renamerRepo renamerRepository) *Service {
	if creatorRepo == nil {
		panic("nil creatorRepository")
	}

	if renamerRepo == nil {
		panic("nil renamerRepository")
	}

	return &Service{creatorRepo: creatorRepo, renamerRepo: renamerRepo}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Contractor, error) {
	contractor, err := NewContractor(input.Name, input.INN)
	if err != nil {
		return Contractor{}, err
	}

	created, err := s.creatorRepo.Save(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("save contractor: %w", err)
	}

	return created, nil
}

func (s *Service) Rename(ctx context.Context, id int64, name string) (Contractor, error) {
	if id <= 0 {
		return Contractor{}, ErrInvalidID
	}

	contractor, err := s.renamerRepo.FindByID(ctx, id)
	if err != nil {
		return Contractor{}, fmt.Errorf("find contractor by id: %w", err)
	}

	err = contractor.Rename(name)
	if err != nil {
		return Contractor{}, fmt.Errorf("rename contractor: %w", err)
	}

	contractor, err = s.renamerRepo.Update(ctx, contractor)
	if err != nil {
		return Contractor{}, fmt.Errorf("update contractor: %w", err)
	}

	return contractor, nil
}
