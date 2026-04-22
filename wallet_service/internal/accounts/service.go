package accounts

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/alvis/wallet_service/internal/httpx"
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
		return nil, httpx.BadRequest("name shoule not be empty")
	}
	return s.store.Create(ctx, name)
}

func (s *Service) GetAccount(ctx context.Context, id uuid.UUID) (*Account, error) {
	acc, err := s.store.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httpx.NotFound("account not found")
		}
		return nil, err
	}
	return acc, nil
}
