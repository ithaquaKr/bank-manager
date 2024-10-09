package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/customer"
	"github.com/ithaquaKr/bank-manager/internal/customer/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MongoRepository Implement of the Repository interface using MongoDB
type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

// NewMongoRepository Create a new instance of the MongoRepository
func NewMongoRepository(client *mongo.Client, dbName string, collName string) customer.Repository {
	return &MongoRepository{
		client: client,
		db:     client.Database(dbName),
		coll:   client.Database(dbName).Collection(collName),
	}
}

// Insert Insert one Customer document into database
func (r *MongoRepository) Insert(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	_, err := r.coll.InsertOne(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("Fail to insert customer: %w", err)
	}

	return customer, nil
}

// InsertMany Insert many Customer document into database
func (r *MongoRepository) InsertMany(ctx context.Context, customers []*models.Customer) error {
	var mongoCustomers []interface{}
	for _, c := range customers {
		mongoCustomers = append(mongoCustomers, c)
	}

	_, err := r.coll.InsertMany(ctx, mongoCustomers)
	if err != nil {
		return fmt.Errorf("Fail to insert list of customer: %w", err)
	}

	return nil
}

// customer.Repository.FindOne retrieves a Customer document by its CustomerID.
func (r *MongoRepository) FindOne(ctx context.Context, customerID uuid.UUID) (*models.Customer, error) {
	filter := bson.M{"customer_id": customerID}
	var customer models.Customer
	err := r.coll.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Customer not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to find customer: %w", err)
	}

	return &customer, nil
}

// FindAll List all Customer in database
func (r *MongoRepository) FindAll(ctx context.Context) ([]*models.Customer, error) {
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Failed to find customer: %w", err)
	}
	defer cursor.Close(ctx)

	var customers []*models.Customer
	for cursor.Next(ctx) {
		var customer models.Customer
		if err := cursor.Decode(&customer); err != nil {
			return nil, fmt.Errorf("Failed to decode customer: %w", err)
		}
		customers = append(customers, &customer)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor err: %w", err)
	}

	return customers, nil
}

// Update Update customer by id
func (r *MongoRepository) Update(ctx context.Context, customerID uuid.UUID, customer *models.Customer) error {
	filter := bson.M{
		"CustomerID": customerID,
	}
	update := bson.M{
		"$set": customer,
	}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Failed to update customer: %w", err)
	}

	return nil
}

// Delete Delete customer by id
func (r *MongoRepository) Delete(ctx context.Context, customerID uuid.UUID) error {
	filter := bson.M{"CustomerID": customerID}
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("Failed to delete customer: %w", err)
	}

	return nil
}
