package customer

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/customer/models"
)

// Customer UseCase interface
type UseCase interface {
	GetById(ctx context.Context, CustomerID uuid.UUID) (*models.Customer, error)
}
