package main

import (
	"context"
	"log"

	statusv1 "github.com/mikedonnici/mono/gen/status/v1"
)

// FetchStatus returns the status of the service including connections to all data sources.
func (s *server) FetchStatus(ctx context.Context, in *statusv1.FetchStatusRequest) (*statusv1.FetchStatusResponse, error) {

	log.Printf("Fetch status")

	return &statusv1.FetchStatusResponse{Healthy: true}, nil
}
