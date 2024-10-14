package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionHistory struct {
	ID              uuid.UUID
	UserId          uuid.UUID
	TransactionType string
	Amount          float64
	BalanceBefore   float64
	BalanceAfter    float64
	Remarks         string
	ReferenceId     string
	CreatedDate     time.Time
	UpdatedDate     time.Time
	Transaction     Transaction `gorm:"foreignKey:ReferenceId"`
}
