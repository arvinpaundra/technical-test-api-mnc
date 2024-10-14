package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber string
	Pin         string
	Balance     float64
	Address     string
	CreatedDate time.Time
	UpdatedDate time.Time
}
