package main

import (
	"context"
	"log"

	userpbv1 "github.com/mikedonnici/mono/internal/user/pb/v1"
)

// FetchUserByID fetches a user by unique id.
func (s *server) FetchUserByID(ctx context.Context, in *userpbv1.FetchUserByIDRequest) (*userpbv1.FetchUserByIDResponse, error) {
	log.Printf("Fetch user...")

	a, err := s.userManager.UserByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	msg := a.Message()
	return &userpbv1.FetchUserByIDResponse{
		User: &msg,
	}, nil
}
