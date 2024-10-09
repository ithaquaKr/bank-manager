package customer

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Customer HTTP Handlers interface
type Handlers interface {
	GetByID() echo.HandlerFunc
}

// Customer handlers
type customerHandlers struct {
	cfg    *config.Config
	logger logger.Logger
	cusUC  UseCase
}

// NewCustomerHandlers Customer handlers constructor
func NewCustomerHandlers(cfg *config.Config, cusUC UseCase, logger logger.Logger) Handlers {
	return &customerHandlers{cfg: cfg, cusUC: cusUC, logger: logger}
}

// GetByID
// @Summary Get customer
// @Description Get customer by id
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Customer
// @Failure 500 {object} httpErrors.RestErr
// @Router /customers/{id} [get]
func (h *customerHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			// LOG RESPONSE ERROR
			return c.JSON(http.StatusInternalServerError)
		}

		customer, err := h.cusUC.GetByID(ctx, customerID)
		if err != nil {
			// LOG RESPONSE ERROR
			return c.JSON()
		}
		return c.JSON(http.StatusOK, customer)
	}
}
