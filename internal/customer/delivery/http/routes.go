package http

import (
	"github.com/ithaquaKr/bank-manager/internal/customer"
	"github.com/labstack/echo/v4"
)

// MapCustomerRoutes Map customer routes http
func MapCustomerRoutes(customerGroup *echo.Group, h customer.Handlers) {
	customerGroup.GET("/:id", h.GetById())
}
