package main

import (
	"github.com/mikedonnici/mono/internal/attribute"
	attributepbv1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"
	"github.com/mikedonnici/mono/internal/status"
	statuspbv1 "github.com/mikedonnici/mono/internal/status/pb/v1"
	"github.com/mikedonnici/mono/internal/user"
	userpbv1 "github.com/mikedonnici/mono/internal/user/pb/v1"
)

type server struct {
	statuspbv1.UnimplementedStatusServiceServer
	attributepbv1.UnimplementedAttributeServiceServer
	userpbv1.UnimplementedUserServiceServer

	attributeManager attribute.Manager
	statusManager    status.Manager
	userManager      user.Manager
}
