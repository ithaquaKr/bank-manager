package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/internal/employee"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	"github.com/labstack/echo/v4"
)

// employeeHandlers Custoemr handler
type employeeHandlers struct {
	cfg        *config.Config
	employeeUC employee.UseCase
	logger     logger.Logger
}

func NewEmployeeHandlers(cfg *config.Config, employeeUC employee.UseCase, logger logger.Logger) employee.Handlers {
	return &employeeHandlers{cfg: cfg, employeeUC: employeeUC, logger: logger}
}

// GetByID
// @Summary Get employee
// @Description Get employee by id
// @Tags Employees
// @Accept  json
// @Produce  json
// @Param id path int true "employee ID"
// @Success 200 {object} models.Employee
// @Router /employees/{id} [get]
func (h *employeeHandlers) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		employeeId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "Invaid employee ID: %w", err))
		}

		employee, err := h.employeeUC.GetById(context.Background(), employeeId)
		if err != nil {
			h.logger.Errorf("ErrorResponseWithLog, Error: %s", err)
			return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error: %w", err))
		}

		return c.JSON(http.StatusOK, employee)
	}
}
