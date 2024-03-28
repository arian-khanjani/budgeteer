package repository

import (
	"budgeteer/internal/model"
	"errors"
)

var (
	ErrConvertExpense = errors.New("couldn't convert data to Expense")
	ErrConvertLoan    = errors.New("couldn't convert data to Loan")
)

func (r *res) ToExpenses() (*[]*model.Expense, error) {
	val, ok := r.value.(*[]*model.Expense)
	if !ok {
		return nil, ErrConvertExpense
	}

	return val, nil
}

func (r *res) ToExpense() (*model.Expense, error) {
	val, ok := r.value.(*model.Expense)
	if !ok {
		return nil, ErrConvertExpense
	}

	return val, nil
}

func (r *res) ToLoans() (*[]*model.Loan, error) {
	val, ok := r.value.(*[]*model.Loan)
	if !ok {
		return nil, ErrConvertLoan
	}

	return val, nil
}

func (r *res) ToLoan() (*model.Loan, error) {
	val, ok := r.value.(*model.Loan)
	if !ok {
		return nil, ErrConvertLoan
	}

	return val, nil
}
