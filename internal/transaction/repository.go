package transaction

import (
	"context"

	"github.com/ithaquaKr/bank-manager/internal/transaction/models"
)

type Repository interface {
	InsertMany(ctx context.Context, transactions []*models.Transaction) error
}
