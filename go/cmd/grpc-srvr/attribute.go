package main

import (
	"context"
	"log"

	attributev1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"
)

// FetchAttribute returns a single attribute
func (s *server) FetchAttribute(ctx context.Context, in *attributev1.FetchAttributeRequest) (*attributev1.FetchAttributeResponse, error) {
	log.Printf("Fetch attribute...")

	a, err := s.attributeManager.AttributeByID(ctx, *in.Id)
	if err != nil {
		return nil, err
	}
	msg := a.Message()
	return &attributev1.FetchAttributeResponse{
		Attribute: &msg,
	}, nil
}
