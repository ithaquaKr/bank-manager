package http

import (
	"github.com/ithaquaKr/bank-manager/internal/employee"
	"github.com/labstack/echo/v4"
)

func MapEmployeeRoutes(employeeGroup *echo.Group, h employee.Handlers) {
	employeeGroup.GET("/:id", h.GetById())
}
