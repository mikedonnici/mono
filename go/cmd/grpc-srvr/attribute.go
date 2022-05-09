package main

import (
	"context"
	"log"

	"github.com/mikedonnici/mono/internal/attribute"
)

// FetchAttribute returns a single attribute
func (s *server) FetchAttribute(ctx context.Context, in *attribute.FetchAttributeRequest) (*attribute.FetchAttributeResponse, error) {
	log.Printf("Fetch attribute...")

	a, err := s.attributeManager.AttributeByID(ctx, *in.Id)
	if err != nil {
		return nil, err
	}
	raw := a.Raw()
	return &attribute.FetchAttributeResponse{
		Attribute: &raw,
	}, nil
}
