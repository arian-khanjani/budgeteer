package model

import "time"

// TODO: give JSON tags

type Loan struct {
	ID               string    `bson:"_id"`
	Timestamp        time.Time `bson:"timestamp"`
	Sum              int64     `bson:"sum"`
	InstallmentCount int       `bson:"installment_count"`
	Status           Status    `bson:"-"` // not stored - calculated upon request
}

type Installment struct {
	ID      string    `bson:"_id"`
	LoanID  string    `bson:"loan_id"`
	Number  int       `bson:"number'"`
	DueDate time.Time `bson:"due_date"`
	Amount  int64     `bson:"amount"`
	Status  Status    `bson:"status'"`
}

type Status int

const (
	ACTIVE Status = iota
	PAYED
)
