package models

import (
	"time"

	"github.com/google/uuid"
)

// Employee Employee models to save into database
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

// Account Accounts has been created embed in Employee model
type Account struct {
	AccountNumber  string    `bson:"account_number"`
	Type           string    `json:"type"`
	InitialDeposit int       `bson:"initial_deposit,omitempty"`
	CreatedAt      time.Time `bson:"created_at"`
}
