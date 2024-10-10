package server

import (
	"github.com/ithaquaKr/bank-manager/docs"
	customerHttpHandler "github.com/ithaquaKr/bank-manager/internal/customer/delivery/http"
	customerRepository "github.com/ithaquaKr/bank-manager/internal/customer/repository"
	customerUsecase "github.com/ithaquaKr/bank-manager/internal/customer/usecase"
	employeeHttpHandler "github.com/ithaquaKr/bank-manager/internal/employee/delivery/http"
	employeeRepository "github.com/ithaquaKr/bank-manager/internal/employee/repository"
	employeeUsecase "github.com/ithaquaKr/bank-manager/internal/employee/usecase"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	customerRepo := customerRepository.NewMongoRepository(s.mongoDb, "bank-manager", "customer")
	employeeRepo := employeeRepository.NewMongoRepository(s.mongoDb, "bank-manager", "employee")

	// Init Usecase
	customerUC := customerUsecase.NewCustomerUsecase(s.cfg, customerRepo, s.appLogger)
	employeeUC := employeeUsecase.NewEmployeeUsecase(s.cfg, employeeRepo, s.appLogger)

	// Init Handler
	customerHandler := customerHttpHandler.NewCustomerHandlers(s.cfg, customerUC, s.appLogger)
	employeeHandler := employeeHttpHandler.NewEmployeeHandlers(s.cfg, employeeUC, s.appLogger)

	v1 := e.Group("/api/v1")
	customerGroup := v1.Group("/customers")
	employeeGroup := v1.Group("/employees")

	customerHttpHandler.MapCustomerRoutes(customerGroup, customerHandler)
	employeeHttpHandler.MapEmployeeRoutes(employeeGroup, employeeHandler)

	docs.SwaggerInfo.Title = "Bank manager REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return nil
}
