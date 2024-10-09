package employee

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	EmployeeId      uuid.UUID `bson:"employee_id"`
	IdentityCard    string    `bson:"identity_card"`
	Fullname        string    `bson:"fullname"`
	DateOfBirth     string    `bson:"date_of_birth"`
	Address         string    `bson:"address"`
	JobRank         string    `bson:"job_rank"`
	Seniority       int       `bson:"seniority"`
	Position        string    `bson:"position"`
	AccountsCreated []Account `bson:"accounts_created,omitempty"`
}

// Add func to generate EmployeeId
type Account struct {
	AccountNumber  string    `bson:"account_number"`
	Type           string    `json:"type"`
	InitialDeposit int       `bson:"initial_deposit,omitempty"`
	CreatedAt      time.Time `bson:"created_at"`
}
