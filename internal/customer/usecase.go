package customer

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer UseCase interface
type UseCase interface {
	GetByID(ctx context.Context, CustomerID *primitive.ObjectID) *Customer
}
