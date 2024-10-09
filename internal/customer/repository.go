package customer

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/customer/models"
)

// Repository Define the interface for interacting with Customer data
type Repository interface {
	Insert(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	InsertMany(ctx context.Context, customers []*models.Customer) error
	FindOne(ctx context.Context, customerID uuid.UUID) (*models.Customer, error)
	FindAll(ctx context.Context) ([]*models.Customer, error)
	Update(ctx context.Context, customerID uuid.UUID, customer *models.Customer) error
	Delete(ctx context.Context, customerID uuid.UUID) error
}
