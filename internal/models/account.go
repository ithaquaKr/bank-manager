package models

import "github.com/google/uuid"

type Account struct {
	AccountID    uuid.UUID `bson:"account_id"`
	Type         string    `bson:"type"`
	Balance      float64   `bson:"balance"`
	CustomerID   uuid.UUID `bson:"customer_id"`
	CreditLimit  *float64  `bson:"credit_limit,omitempty"`
	InterestRate *float64  `bson:"interest_rate,omitempty"`
	MinBalance   *float64  `bson:"min_balance,omitempty"`
}
