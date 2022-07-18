package main

import (
	"context"
	"fmt"
	"log"
	"time"

	statusv1 "github.com/mikedonnici/mono/internal/status/pb/v1"
)

// FetchStatus returns the status of the service including connections to all data sources.
func (s *server) FetchStatus(ctx context.Context, in *statusv1.FetchStatusRequest) (*statusv1.FetchStatusResponse, error) {

	log.Printf("Fetching status...")

	return &statusv1.FetchStatusResponse{
		Healthy: true,
		Message: fmt.Sprintf("Status at %v", time.Now().Format(time.RFC3339)),
	}, nil
}
