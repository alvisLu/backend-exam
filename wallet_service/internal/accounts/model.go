package accounts

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID        uuid.UUID       `gorm:"column:id;type:uuid;primaryKey"`
	Name      string          `gorm:"column:name;type:text;not null"`
	Balance   decimal.Decimal `gorm:"column:balance;type:numeric(20,8);not null"`
	CreatedAt time.Time       `gorm:"column:created_at;type:timestamptz;not null"`
}

func (Account) TableName() string {
	return "accounts"
}
