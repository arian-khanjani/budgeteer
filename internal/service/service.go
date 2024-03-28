package service

import (
	"budgeteer/internal/model"
	"budgeteer/internal/repository"
	"context"
	"errors"
)

type Service struct {
	repo repository.Repository
}

// New creates a metadata service controller.
func New(repo repository.Repository) *Service {
	return &Service{repo}
}

// GetExpenses returns movie metadata by id.
func (s *Service) GetExpenses(ctx context.Context, id string) ([]*model.Expense, error) {
	res, err := s.repo.GetExpenses()
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, repository.ErrNotFound
	}

	return res, nil
}
