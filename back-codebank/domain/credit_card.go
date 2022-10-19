package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID              string
	Name            string
	Numner          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Balance         float64
	Limit           float64
	CreatedAt       time.Time
}

func NewCreditCard() *CreditCard { //Pointer
	creditCard := &CreditCard{}
	creditCard.ID = uuid.NewV4().String()
	creditCard.CreatedAt = time.Now()
	return creditCard
}
