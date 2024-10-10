package models

import "github.com/google/uuid"

type Transaction struct {
	TransactionId uuid.UUID `bson:"transaction_id"`
	CustomerID    uuid.UUID `bson:"customer_id"`
	EmployeeID    uuid.UUID `bson:"employee_id"`
	AccountNumber string    `bson:"account_number"`
	Type          string    `bson:"type"` // "deposit", "withdraw", "payment", ...
	Amount        float64   `bson:"amount"`
	CreatedAt     string    `bson:"created_at"` // TODO: Handle this datetime field
}
