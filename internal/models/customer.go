package models

import "github.com/google/uuid"

type Customer struct {
	IdentityCard string    `bson:"identity_card"`
	CustomerID   uuid.UUID `bson:"customer_id"`
	Name         string    `bson:"name"`
	DateOfBirth  string    `bson:"date_of_birth"`
	Address      string    `bson:"address"`
	Accounts     []string  `bson:"accounts"`
}
