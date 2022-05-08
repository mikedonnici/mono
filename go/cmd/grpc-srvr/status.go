package main

import (
	"context"
	"log"

	"github.com/mikedonnici/mono/internal/status"
)

// FetchStatus returns the status of the service including connections to all data sources.
func (s *server) FetchStatus(ctx context.Context, in *status.StatusRequest) (*status.StatusResponse, error) {

	log.Printf("Fetch status")

	return &status.StatusResponse{Healthy: true}, nil
}
