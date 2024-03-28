package mongodb

import "budgeteer/internal/model"

type Repository struct {
}

func New() (*Repository, error) {
	// TODO: os.GetEnv()...
	return &Repository{}, nil
}

func (r *Repository) GetExpenses() ([]*model.Expense, error) {

	return nil, nil
}
