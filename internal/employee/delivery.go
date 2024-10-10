package employee

import "github.com/labstack/echo/v4"

// Handlers HTTP Handlers interface for Employee
type Handlers interface {
	GetById() echo.HandlerFunc
}
