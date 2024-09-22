// Package cmd implements the cobra CLI for Bank-manager server.
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ithaquaKr/bank-manager/config"
	"github.com/ithaquaKr/bank-manager/pkg/logger"
	"github.com/ithaquaKr/bank-manager/pkg/server"
	"github.com/spf13/cobra"
)

// -----------------------------------Global constant BEGIN----------------------------------------.
const (
	// greetingBanner is the greeting banner.
	greetingBanner = `
___________________________________________________________________________________________

██╗  ██╗███████╗██╗     ██╗      ██████╗
██║  ██║██╔════╝██║     ██║     ██╔═══██╗
███████║█████╗  ██║     ██║     ██║   ██║
██╔══██║██╔══╝  ██║     ██║     ██║   ██║
██║  ██║███████╗███████╗███████╗╚██████╔╝
╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝ ╚═════╝

%s
___________________________________________________________________________________________

`
	// byeBanner is the bye banner.
	byeBanner = `
██████╗ ██╗   ██╗███████╗
██╔══██╗╚██╗ ██╔╝██╔════╝
██████╔╝ ╚████╔╝ █████╗
██╔══██╗  ╚██╔╝  ██╔══╝
██████╔╝   ██║   ███████╗
╚═════╝    ╚═╝   ╚══════╝
`
)

// -----------------------------------Global constant END------------------------------------------

// -----------------------------------Command Line Config BEGIN------------------------------------
var (
	flags struct {
		// TODO: Implement flags
	}

	rootCmd = &cobra.Command{
		Use:   "bankm",
		Short: "Bank manager is a api server for manage account, client, transfer in Bank",
		Run: func(_ *cobra.Command, _ []string) {
			start()
			fmt.Printf("%s", byeBanner)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// TODO: Init flags for command
}

// -----------------------------------Command Line Config END--------------------------------------

func checkPort(port int) error {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return err
	}
	return l.Close()
}

func start() {
	config, err := config.InitConfig(".", "config")
	if err != nil {
		// TODO: Convert to using global log format
		log.Fatalf("Cannot get env variable, err: %s", err)
	}
	apiLogger := logger.NewApiLogger(config)
	apiLogger.InitLogger()
	apiLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s",
		config.App.AppVersion,
		config.Logger.Level,
		config.App.Mode,
	)
	// Strangely, when the port is unavailable, echo server would return OK response for /healthz
	// and then complain unable to bind port. Thus we cannot rely on checking /healthz. As a
	// workaround, we check whether the port is available here.
	if err := checkPort(8000); err != nil {
		// slog.Error(fmt.Sprintf("server port %d is not available", flags.port), log.BBError(err))
		fmt.Errorf("server port is not available", err)
		return
	}

	var s *server.Server
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	// Trigger graceful shutdown on SIGINT or SIGTERM,
	// which is taken as the graceful shutdown signal for many systems, eg., Kubernetes, Gunicorn.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Println(fmt.Sprintf("%s received.", sig.String()))
		if s != nil {
			_ = s.Shutdown(ctx)
		}
		cancel()
	}()

	s = server.NewServer(ctx, config, apiLogger)

	fmt.Printf(greetingBanner, fmt.Sprintf("Version %s has started on port %d 🚀", config.App.AppVersion, config.App.Port))

	// Execute the program.
	if err := s.Run(ctx, config.App.Port); err != nil {
		if err != http.ErrServerClosed {
			// TODO: Fix that log
			log.Fatal(err.Error())
			_ = s.Shutdown(ctx)
			cancel()
		}
	}

	// Wait for CTRL-C
	<-ctx.Done()
}
