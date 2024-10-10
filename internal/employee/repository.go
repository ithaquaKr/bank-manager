package employee

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/employee/models"
)

// Repository Interface for interacting with employee data
type Repository interface {
	Insert(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	InsertMany(ctx context.Context, employees []*models.Employee) error
	FindOne(ctx context.Context, employeeID uuid.UUID) (*models.Employee, error)
	FindAll(ctx context.Context) ([]*models.Employee, error)
	Update(ctx context.Context, employeeID uuid.UUID, employee *models.Employee) error
	Delete(ctx context.Context, employeeID uuid.UUID) error
}
