package transfers

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/alvis/wallet_service/internal/accounts"
	"github.com/alvis/wallet_service/internal/httpx"
)

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

type Result struct {
	From   *accounts.Account
	To     *accounts.Account
	Amount decimal.Decimal
}

func (s *Service) Transfer(ctx context.Context, fromIDStr, toIDStr, amountStr string) (*Result, error) {
	fromID, err := uuid.Parse(fromIDStr)
	if err != nil {
		return nil, httpx.BadRequest("invalid from_id")
	}
	toID, err := uuid.Parse(toIDStr)
	if err != nil {
		return nil, httpx.BadRequest("invalid to_id")
	}

	if fromID == toID {
		return nil, httpx.BadRequest("from_id and to_id must be different")
	}

	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return nil, httpx.BadRequest("invalid amount")
	}
	if !amount.IsPositive() {
		return nil, httpx.BadRequest("amount must be positive")
	}
	if amount.Exponent() < -8 {
		return nil, httpx.BadRequest("amount exceeds 8 decimal places")
	}

	from, to, err := s.store.ExecuteTransfer(ctx, fromID, toID, amount)
	if err != nil {
		return nil, err
	}
	return &Result{From: from, To: to, Amount: amount}, nil
}
