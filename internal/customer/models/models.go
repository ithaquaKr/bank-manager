package models

import "github.com/google/uuid"

type Customer struct {
	CustomerId   uuid.UUID `bson:"customer_id"`
	IdentityCard string    `bson:"identity_card"`
	Name         string    `bson:"name"`
	DateOfBirth  string    `bson:"date_of_birth"` // TODO: Handle this datatime field
	Address      string    `bson:"address"`
	Accounts     []Account `bson:"accounts"`
}

type Account struct {
	AccountNumber string `bson:"account_number"`
	Fullname      string `bson:"fullname"`
	AccountType   string `bson:"account_type"` // "debit", "credit"
	// Debit account type
	Balance             *float64 `bson:"balance,omitempty"`
	MonthlyInterestRate *float64 `bson:"monthly_interest_rate,omitempty"`
	MinBalance          *float64 `bson:"min_balance,omitempty"`
	// Credit account type
	CreditLimit        *float64 `bson:"credit_limit,omitempty"`
	OutstandingBalance *float64 `bson:"outstading_balance"`

	Transactions []Transaction `bson:"transaction,omitempty"`
}

type Transaction struct {
	TransactionId uuid.UUID `bson:"transaction_id"`
	Amount        float64   `bson:"amount"`
	CreatedAt     string    `bson:"created_at"` // TODO: Handle this datetime field too
}
