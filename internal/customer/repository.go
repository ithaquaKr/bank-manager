package customer

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// customer.Repository defines the interface for interacting with Customer data
type Repository interface {
	Insert(ctx context.Context, customer *Customer) (*Customer, error)
	FindOne(ctx context.Context, customerID primitive.ObjectID) (*Customer, error)
	FindAll(ctx context.Context) ([]*Customer, error)
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, customerID primitive.ObjectID) error
}

// customer.MongoRepository implements the customer.Repository interface using MongoDB
type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

// customer.NewMongoRepository creates a new instance of the customer.MongoRepository
func NewMongoRepository(client *mongo.Client, dbName string, collName string) *MongoRepository {
	return &MongoRepository{
		client: client,
		db:     client.Database(dbName),
		coll:   client.Database(dbName).Collection(collName),
	}
}

// customer.Repository.Insert insert a new Customer document in to database.
func (r *MongoRepository) Insert(ctx context.Context, customer *Customer) (*mongo.InsertOneResult, error) {
	result, err := r.coll.InsertOne(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("customer.MongoRepository.Insert.InsertOne: %w", err)
	}

	return result, nil
}

// customer.Repository.FindOne retrieves a Customer document by its CustomerID.
func (r *MongoRepository) FindOne(ctx context.Context, customerID primitive.ObjectID) (*Customer, error) {
	filter := bson.M{"CustomerID": customerID}
	var customer Customer
	err := r.coll.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Customer not found: %w", err)
		}
		return nil, fmt.Errorf("Failed to find customer: %w", err)
	}

	return &customer, nil
}

// customer.Repository.FindAll retrieves all Customer documents.
func (r *MongoRepository) FindAll(ctx context.Context) ([]*Customer, error) {
	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Failed to find customer: %w", err)
	}
	defer cursor.Close(ctx)

	var customers []*Customer
	for cursor.Next(ctx) {
		var customer Customer
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

// customer.Repository.Update
func (r *MongoRepository) Update(ctx context.Context, customer *Customer) error {
	filter := bson.M{
		"CustomerID": customer.CustomerID,
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

// customer.Repository.Delete
func (r *MongoRepository) Delete(ctx context.Context, customerID *primitive.ObjectID) error {
	filter := bson.M{"CustomerID": customerID}
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("Failed to delete customer: %w", err)
	}

	return nil
}
