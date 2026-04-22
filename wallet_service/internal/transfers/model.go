package transfers

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	ID        uuid.UUID       `gorm:"column:id;type:uuid;primaryKey"`
	FromID    uuid.UUID       `gorm:"column:from_id;type:uuid;not null"`
	ToID      uuid.UUID       `gorm:"column:to_id;type:uuid;not null"`
	Amount    decimal.Decimal `gorm:"column:amount;type:numeric(20,8);not null"`
	CreatedAt time.Time       `gorm:"column:created_at;type:timestamptz;not null"`
}

func (Transfer) TableName() string { return "transfers" }
