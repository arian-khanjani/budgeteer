package memory

import (
	"budgeteer/internal/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Expense
}

func New() (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetExpenses() ([]*model.Expense, error) {
	r.RLock()
	defer r.RUnlock()

	return nil, nil
}
