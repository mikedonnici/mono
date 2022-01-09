package main

import (
	"github.com/mikedonnici/mrv/internal/attribute"
	"github.com/mikedonnici/mrv/internal/statuspb"
)

type server struct {
	statuspb.UnimplementedStatusServiceServer
	attribute.UnimplementedAttributeServiceServer

	attributeManager attribute.Manager
}

func (s *server) attachDataManagers() error {
	s.attributeManager = attribute.Manager{}
	return nil
}
