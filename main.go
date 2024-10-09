// Package main is the main binary for Bank-manager service.
package main

import (
	"os"

	cmd "github.com/ithaquaKr/bank-manager/cmd/server"
)

// @title Bank-manager Swagger API
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name IthaquaKr
// @contact.url https://www.github.com/ithaquaKr
// @contact.email ithadev.nguyen@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
