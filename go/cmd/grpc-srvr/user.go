package main

import (
	"context"
	"errors"
	"log"

	"github.com/mikedonnici/mono/internal/user"
	userpbv1 "github.com/mikedonnici/mono/internal/user/pb/v1"
)

// FetchUser fetches a user by unique id or email.
func (s *server) FetchUser(ctx context.Context, in *userpbv1.FetchUserRequest) (*userpbv1.FetchUserResponse, error) {
	log.Printf("Fetch user...")

	var u *user.User
	var err error

	// id should take precedence if both are specified
	switch {
	case in.Id != nil && *in.Id != 0:
		u, err = s.userManager.UserByID(ctx, *in.Id)
	case in.Email != nil && *in.Email != "":
		u, err = s.userManager.UserByEmail(ctx, *in.Email)
	default:
		err = errors.New("both id and email are empty")
	}
	if err != nil {
		return nil, err
	}
	msg := u.Message()
	return &userpbv1.FetchUserResponse{
		User: &msg,
	}, nil
}
