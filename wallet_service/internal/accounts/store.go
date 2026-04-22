package accounts

import (
	"context"

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
