package memory

import (
	"budgeteer/internal/model"
	"context"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Expense
}

func New(ctx context.Context) (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetExpenses() ([]*model.Expense, error) {
	r.RLock()
	defer r.RUnlock()

	return nil, nil
}
