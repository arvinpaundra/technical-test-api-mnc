package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	AccessToken  string
	RefreshToken string
	CreatedDate  time.Time
	UpdatedDate  time.Time
}
