package main

import (
	"github.com/mikedonnici/mono/internal/attribute"
	"github.com/mikedonnici/mono/internal/status"
)

type server struct {
	status.UnimplementedStatusServiceServer
	attribute.UnimplementedAttributeServiceServer

	attributeManager attribute.Manager
	statusManager    status.Manager
}
