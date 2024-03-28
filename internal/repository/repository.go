package repository

import (
	"errors"
)

// TODO: hardcode queries!

type Repository interface {
	New() (Repository, error)
	Create(interface{}) error
	List() (*[]*res, error)
	Get(string) (*res, error)
	Update(string, interface{}) (*res, error)
}

type res struct {
	value interface{}
}

var (
	ErrNotFound = errors.New("document not found")
)
