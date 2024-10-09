package customer

import "github.com/google/uuid"

type Customer struct {
	CustomerId   string   `bson:"customer_id"`
	IdentityCard string   `bson:"identity_card"`
	Name         string   `bson:"name"`
	DateOfBirth  string   `bson:"date_of_birth"`
	Address      string   `bson:"address"`
	Accounts     []string `bson:"accounts"`
}

type Account struct {
	AccountId    uuid.UUID `bson:"account_id"`
	Name         string    `bson:"name"`
	Type         string    `bson:"type"`
	Balance      float64   `bson:"balance"`
	CustomerId   uuid.UUID `bson:"customer_id"`
	CreditLimit  *float64  `bson:"credit_limit,omitempty"`
	InterestRate *float64  `bson:"interest_rate,omitempty"`
	MinBalance   *float64  `bson:"min_balance,omitempty"`
}

type Transaction struct {
	TransactionId uuid.UUID `bson:"id"`
	Amount        float64   `bson:"amount"`
	AccountID     uuid.UUID `bson:"account_id"`
	Type          string    `bson:"type"` // "deposit", "withdraw", "payment"
	EmployeeID    uuid.UUID `bson:"employee_id"`
	Date          string    `bson:"date"`
	Description   string    `bson:"description"`
	SourceAccount *string   `bson:"source_account,omitempty"`
}
