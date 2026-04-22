package accounts

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(ctx context.Context, name string) (*Account, error) {
	acc := &Account{Name: name}
	if err := s.db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Omit("ID", "Balance", "CreatedAt").
		Create(acc).Error; err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *Store) GetByID(ctx context.Context, id uuid.UUID) (*Account, error) {
	var acc Account
	if err := s.db.WithContext(ctx).First(&acc, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &acc, nil
}
