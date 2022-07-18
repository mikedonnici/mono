package main

import (
	"github.com/mikedonnici/mono/internal/attribute"
	attributev1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"
	"github.com/mikedonnici/mono/internal/status"
	statusv1 "github.com/mikedonnici/mono/internal/status/pb/v1"
)

type server struct {
	statusv1.UnimplementedStatusServiceServer
	attributev1.UnimplementedAttributeServiceServer

	attributeManager attribute.Manager
	statusManager    status.Manager
}
