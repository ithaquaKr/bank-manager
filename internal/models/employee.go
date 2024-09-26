package models

import "github.com/google/uuid"

type Employee struct {
	EmployeeID   uuid.UUID `bson:"employee_id"`
	IdentityCard string    `bson:"identity_card"`
	Name         string    `bson:"name"`
	DateOfBirth  string    `bson:"date_of_birth"`
	Address      string    `bson:"address"`
	JobRank      string    `bson:"job_rank"`
	Seniority    int       `bson:"seniority"`
	JobPosition  string    `bson:"job_position"`
	IsSalesStaff bool      `bson:"is_sales_staff"`
}
