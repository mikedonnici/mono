package main

import (
	"context"
	"log"

	"github.com/mikedonnici/mrv/internal/status"
)

// FetchStatus returns the status of the server
func (s *server) FetchStatus(ctx context.Context, in *status.StatusRequest) (*status.StatusResponse, error) {
	log.Printf("Check status")
	return &status.StatusResponse{Healthy: true}, nil
}
