package main

import (
	"log"

	"github.com/hc12r/sentence-analyzer-vm/internal/server"
)

func main() {
	// Setup and run the server using the internal server package
	if err := server.SetupAndRun(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
