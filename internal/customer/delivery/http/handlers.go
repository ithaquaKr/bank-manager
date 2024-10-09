package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/internal/customer"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	"github.com/labstack/echo/v4"
)

// customerHandlers Custoemr handler
type customerHandlers struct {
	cfg        *config.Config
	customerUC customer.UseCase
	logger     logger.Logger
}

func NewCustomerHandlers(cfg *config.Config, customerUC customer.UseCase, logger logger.Logger) customer.Handlers {
	return &customerHandlers{cfg: cfg, customerUC: customerUC, logger: logger}
}

// GetByID
// @Summary Get customer
// @Description Get customer by id
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} models.Customer
// @Router /customers/{id} [get]
func (h *customerHandlers) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "Invaid customer ID: %w", err))
		}

		customer, err := h.customerUC.GetById(context.Background(), customerId)
		if err != nil {
			h.logger.Errorf("ErrorResponseWithLog, Error: %s", err)
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error: %w", err))
		}

		return c.JSON(http.StatusOK, customer)
	}
}
