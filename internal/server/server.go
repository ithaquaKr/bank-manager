package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	mymongo "github.com/ithaquaKr/bank-manager/pkg/store/mongo"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var gracefulShutdownPeriod = 10 * time.Second

type Server struct {
	echoServer *echo.Echo
	cfg        *config.Config
	appLogger  logger.Logger
	mongoDb    *mongo.Client
}

// NewServer create a new server.
func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	s := &Server{cfg: cfg}

	// Init App logger
	appLogger := logger.NewApiLogger(cfg)
	s.appLogger = appLogger
	s.appLogger.InitLogger()

	// Display configs
	s.appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.App.AppVersion,
		cfg.Logger.Level,
		cfg.App.Mode,
	)

	// Configure Echo server
	s.echoServer = echo.New()
	s.echoServer.HideBanner = true
	s.echoServer.HidePort = true

	// Config MongoDB Client
	mongo, err := mymongo.NewMongoClient(cfg)
	if err != nil {
		s.appLogger.Fatalf("Cannot connect to MongoDB: %v", err)
	}
	s.mongoDb = mongo
	s.appLogger.Info("MongoDB Connected.")

	return s, nil
}

// Run will start the Server
func (s *Server) Run(ctx context.Context, port int) error {
	address := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr: address,
	}

	go func() {
		if err := s.echoServer.StartServer(server); err != nil {
			s.appLogger.Errorf("Error while starting Server: %s", err)
		}
	}()

	// Map server handler
	if err := s.MapHandlers(s.echoServer); err != nil {
		return err
	}

	return nil
}

// Shutdown will shut down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.appLogger.Info("Stoping Server ...")

	ctx, cancel := context.WithTimeout(ctx, gracefulShutdownPeriod)
	defer cancel()

	// Close MongoDB Connection
	if err := s.mongoDb.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	// Gracefully shutdown Echo Server
	if err := s.echoServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown Echo server: %w", err)
	}

	return nil
}
