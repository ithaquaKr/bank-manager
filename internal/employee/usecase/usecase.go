package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/internal/employee"
	"github.com/ithaquaKr/bank-manager/internal/employee/models"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
)

type employeeUC struct {
	cfg          *config.Config
	employeeRepo employee.Repository
	logger       logger.Logger
}

func NewEmployeeUsecase(cfg *config.Config, employeeRepo employee.Repository, logger logger.Logger) employee.UseCase {
	return &employeeUC{cfg: cfg, employeeRepo: employeeRepo, logger: logger}
}

func (u *employeeUC) GetById(ctx context.Context, employeeID uuid.UUID) (*models.Employee, error) {
	return u.employeeRepo.FindOne(ctx, employeeID)
}
