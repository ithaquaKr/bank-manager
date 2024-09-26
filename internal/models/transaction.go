package models

import "github.com/google/uuid"

type Transaction struct {
	TransactionID uuid.UUID `bson:"transaction_id"`
	Amount        float64   `bson:"amount"`
	AccountID     uuid.UUID `bson:"account_id"`
	Type          string    `bson:"type"` // "deposit", "withdraw", "payment"
	EmployeeID    uuid.UUID `bson:"employee_id"`
	Date          string    `bson:"date"`
	Description   string    `bson:"description"`
	SourceAccount *string   `bson:"source_account,omitempty"`
}
