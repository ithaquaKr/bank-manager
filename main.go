// Package main is the main binary for Bank-manager service.
package main

import (
	"os"

	cmd "github.com/ithaquaKr/bank-manager/cmd/server"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
