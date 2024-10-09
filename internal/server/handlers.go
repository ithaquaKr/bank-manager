package server

import (
	"github.com/ithaquaKr/bank-manager/internal/customer"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	cusRepo := customer.NewMongoRepository(s.mongoDb, "testdb", "testcoll")
	if cusRepo != nil {
		return nil
	}

	return nil
}
