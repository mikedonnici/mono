package main

import (
	attributev1 "github.com/mikedonnici/mono/gen/attribute/v1"
	statusv1 "github.com/mikedonnici/mono/gen/status/v1"
	"github.com/mikedonnici/mono/internal/attribute"
	"github.com/mikedonnici/mono/internal/status"
)

type server struct {
	statusv1.UnimplementedStatusServiceServer
	attributev1.UnimplementedAttributeServiceServer

	attributeManager attribute.Manager
	statusManager    status.Manager
}
