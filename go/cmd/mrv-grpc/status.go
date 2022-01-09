package main

import (
	"context"
	"log"

	"github.com/mikedonnici/mrv/internal/statuspb"
)

// FetchStatus returns the status of the server
func (s *server) FetchStatus(ctx context.Context, in *statuspb.StatusRequest) (*statuspb.StatusResponse, error) {
	log.Printf("Check status")
	return &statuspb.StatusResponse{Healthy: true}, nil
}
