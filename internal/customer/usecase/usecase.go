package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/internal/customer"
	"github.com/ithaquaKr/bank-manager/internal/customer/models"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
)

type customerUC struct {
	cfg          *config.Config
	customerRepo customer.Repository
	logger       logger.Logger
}

func NewCustomerUsecase(cfg *config.Config, customerRepo customer.Repository, logger logger.Logger) customer.UseCase {
	return &customerUC{cfg: cfg, customerRepo: customerRepo, logger: logger}
}

func (u *customerUC) GetById(ctx context.Context, customerID uuid.UUID) (*models.Customer, error) {
	return u.customerRepo.FindOne(ctx, customerID)
}
