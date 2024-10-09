package customer

import "github.com/labstack/echo/v4"

// Handlers HTTP Handlers interface
type Handlers interface {
	GetById() echo.HandlerFunc
}
