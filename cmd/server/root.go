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
	"github.com/ithaquaKr/bank-manager/internal/server"
	"github.com/spf13/cobra"
)

// -----------------------------------Command Line Config BEGIN------------------------------------
var (
	rootCmd = &cobra.Command{
		Use:   "bank-manager",
		Short: "Bank manager is a api server for manage account, client, transfer in Bank",
		Run: func(_ *cobra.Command, _ []string) {
			start()
			fmt.Println("End....")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
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

	s, err = server.NewServer(ctx, config)
	if err != nil {
		fmt.Errorf("Cannot create Server", err)
	}

	fmt.Printf("Application starting..: %s", fmt.Sprintf("Version %s has started on port %d 🚀", config.App.AppVersion, config.App.Port))

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
