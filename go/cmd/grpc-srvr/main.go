package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mikedonnici/mono/internal/attribute"
	"github.com/mikedonnici/mono/internal/status"
)

func main() {

	var cfg config
	if err := cfg.Set(); err != nil {
		log.Fatal(err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {

	var (
		port = flag.Int("port", cfg.grpcPort, "The server port")
	)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Server
	var srvr server
	s := grpc.NewServer()

	// Register the GRPC services
	status.RegisterStatusServiceServer(s, &srvr)
	attribute.RegisterAttributeServiceServer(s, &srvr)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Adds data managers to the server for convenience
	if err := srvr.attachDataManagers(); err != nil {
		return fmt.Errorf("could not attach data managers: %w", err)
	}

	log.Printf("grpc-srvr listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("could not listen: %w", err)
	}
	return nil
}
