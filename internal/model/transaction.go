package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID
	UserId      uuid.UUID
	TargetUser  uuid.UUID
	Amount      float64
	Remarks     string
	Category    string
	Status      string
	CreatedDate time.Time
	UpdatedDate time.Time
}
