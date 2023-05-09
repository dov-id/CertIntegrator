package main

import (
	"os"

	"github.com/dov-id/CertIntegrator/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
