package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/employee"
	"github.com/ithaquaKr/bank-manager/internal/employee/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MongoRepository Implement of the Repository interface using MongoDB
type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName string, collName string) employee.Repository {
	return &MongoRepository{
		client: client,
		db:     client.Database(dbName),
		coll:   client.Database(dbName).Collection(collName),
	}
}

func (r *MongoRepository) Insert(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	_, err := r.coll.InsertOne(ctx, employee)
	if err != nil {
		return nil, fmt.Errorf("Fail to insert employee", err)
	}

	return employee, nil
}

func (r *MongoRepository) InsertMany(ctx context.Context, employees []*models.Employee) error {
	var mgEmployees []interface{}
	for _, c := range employees {
		mgEmployees = append(mgEmployees, c)
	}

	_, err := r.coll.InsertMany(ctx, mgEmployees)
	if err != nil {
		return fmt.Errorf("Fail to insert list of employee: %w", err)
	}

	return nil
}

func (r *MongoRepository) FindOne(ctx context.Context, employeeID uuid.UUID) (*models.Employee, error) {
	filter := bson.M{"employee_id": employeeID.String()}
	var employee models.Employee
	err := r.coll.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Employee not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to find employee: %w", err)
	}

	return &employee, nil
}

func (r *MongoRepository) FindAll(ctx context.Context) ([]*models.Employee, error) {
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Failed to find employees: %w", err)
	}
	defer cursor.Close(ctx)

	var employees []*models.Employee
	for cursor.Next(ctx) {
		var employee models.Employee
		if err := cursor.Decode(&employee); err != nil {
			return nil, fmt.Errorf("Failed to decode employee", err)
		}
		employees = append(employees, &employee)

	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("Cursor error: %w", err)
	}

	return employees, nil
}

func (r *MongoRepository) Update(ctx context.Context, employeeID uuid.UUID, employee *models.Employee) error {
	filter := bson.M{
		"employee_id": employeeID.String(),
	}
	update := bson.M{
		"$set": employee,
	}

	_, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Failed to update employee: %w", err)
	}

	return nil
}

func (r *MongoRepository) Delete(ctx context.Context, employeeID uuid.UUID) error {
	filter := bson.M{
		"employee_id": employeeID.String(),
	}

	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("Failed to delete employee", err)
	}

	return nil
}
