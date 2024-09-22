package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	"github.com/labstack/echo/v4"
)

var gracefulShutdownPeriod = 10 * time.Second

type Server struct {
	echoServer *echo.Echo
	cfg        *config.Config
	logger     logger.Logger
}

// NewServer create a new server.
func NewServer(ctx context.Context, cfg *config.Config, logger logger.Logger) *Server {
	return &Server{
		cfg:        cfg,
		echoServer: echo.New(),
		logger:     logger,
	}
}

// Run will start the Server
func (s *Server) Run(ctx context.Context, port int) error {
	address := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr: address,
	}

	s.echoServer.HideBanner = true
	s.echoServer.HidePort = true

	go func() {
		if err := s.echoServer.StartServer(server); err != nil {
			s.logger.Errorf("Error while starting Server: %s", err)
		}
	}()

	return nil
}

// Shutdown will shut down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Stoping Server ...")

	ctx, cancel := context.WithTimeout(ctx, gracefulShutdownPeriod)
	defer cancel()

	if s.echoServer != nil {
		if err := s.echoServer.Shutdown(ctx); err != nil {
			s.echoServer.Logger.Fatal(err)
		}
	}

	return nil
}
