package accounts

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrSameAccount       = errors.New("from and to are the same account")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInvalidName       = errors.New("invalid name")
)

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateAccount(ctx context.Context, name string) (*Account, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}
	return s.store.Create(ctx, name)
}
