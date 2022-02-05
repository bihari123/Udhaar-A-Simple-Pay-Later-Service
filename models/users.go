package models

import (
	"github.com/google/uuid"
)

type User struct {
	Name          string
	Email         string
	CreditLimit   float64
	Dues          float64
	TransactionID []uuid.UUID
}


