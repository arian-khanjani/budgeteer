package expenseservice

import "budgeteer/internal/model"

type Repository interface {
	GetExpense(id string) (*[]*model.Expense, error)
	Getloan(id string) (*[]*model.Loan, error)
	GetAllExpenses(id string) (*[]*model.Expense, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {

	return Service{repo: repo}
}
