package server

import (
	"github.com/ithaquaKr/bank-manager/docs"
	customerHttpHandler "github.com/ithaquaKr/bank-manager/internal/customer/delivery/http"
	customerRepository "github.com/ithaquaKr/bank-manager/internal/customer/repository"
	customerUsecase "github.com/ithaquaKr/bank-manager/internal/customer/usecase"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	customerRepo := customerRepository.NewMongoRepository(s.mongoDb, "bank-manager", "customer")

	// Init Usecase
	customerUC := customerUsecase.NewCustomerUsecase(s.cfg, customerRepo, s.appLogger)

	// Init Handler
	customerHandler := customerHttpHandler.NewCustomerHandlers(s.cfg, customerUC, s.appLogger)

	v1 := e.Group("/api/v1")
	customerGroup := v1.Group("/customers")

	customerHttpHandler.MapCustomerRoutes(customerGroup, customerHandler)

	docs.SwaggerInfo.Title = "Bank manager REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return nil
}
