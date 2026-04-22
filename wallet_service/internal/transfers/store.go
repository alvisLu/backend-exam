package transfers

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/alvis/wallet_service/internal/accounts"
	"github.com/alvis/wallet_service/internal/httpx"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func sortPair(a, b uuid.UUID) (lo, hi uuid.UUID) {
	if a.String() < b.String() {
		return a, b
	}
	return b, a
}

func (s *Store) ExecuteTransfer(
	ctx context.Context,
	fromID, toID uuid.UUID,
	amount decimal.Decimal,
) (from, to *accounts.Account, err error) {
	var fromAcc, toAcc accounts.Account

	txErr := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		lo, hi := sortPair(fromID, toID)

		var loAcc, hiAcc accounts.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&loAcc, "id = ?", lo).Error; err != nil {
			return httpx.NotFound("account not found")
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&hiAcc, "id = ?", hi).Error; err != nil {
			return httpx.NotFound("account not found")
		}

		if loAcc.ID == fromID {
			fromAcc, toAcc = loAcc, hiAcc
		} else {
			fromAcc, toAcc = hiAcc, loAcc
		}

		if fromAcc.Balance.LessThan(amount) {
			return httpx.UnprocessableEntity("insufficient funds")
		}

		if err := tx.Model(&accounts.Account{}).Where("id = ?", fromID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Model(&accounts.Account{}).Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		fromAcc.Balance = fromAcc.Balance.Sub(amount)
		toAcc.Balance = toAcc.Balance.Add(amount)

		t := Transfer{FromID: fromID, ToID: toID, Amount: amount}
		return tx.Clauses(clause.Returning{}).
			Omit("ID", "CreatedAt").
			Create(&t).Error
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if txErr != nil {
		return nil, nil, txErr
	}
	return &fromAcc, &toAcc, nil
}
