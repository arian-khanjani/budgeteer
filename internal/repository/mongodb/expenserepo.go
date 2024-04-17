package mongodb

import (
	"budgeteer/internal/model"
	"errors"
)

var (
	ErrGetExpense = errors.New("couldn't Get data to Expense")
	ErrGetLoan    = errors.New("couldn't Get data to Loan")
)

func (r *Repository) GetExpense(id string) ([]*model.Expense, error) {

	return  []*model.Expense, ErrGetExpense

}

func (r *Repository) Getloan(id string) ([]*model.Loan, error) {

	return []*model.Loan, ErrGetLoan

}

func (r *Repository) GetAllExpenses(id string) ([]*model.Expense, error) {

	return []*model.Expense, ErrGetExpense

}
