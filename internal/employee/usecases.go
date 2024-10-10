package employee

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/internal/employee/models"
)

type UseCase interface {
	GetById(ctx context.Context, EmployeeID uuid.UUID) (*models.Employee, error)
}
