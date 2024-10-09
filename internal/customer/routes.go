package customer

import (
	"github.com/labstack/echo/v4"
)

// Map Customer routes
func MapCustomerRoutes(customerGroup *echo.Group, h Handlers) {
	customerGroup.GET("/:id", h.GetByID())
}
