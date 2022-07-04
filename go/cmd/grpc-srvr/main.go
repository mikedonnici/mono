package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	statuspbv1 "github.com/mikedonnici/mono/internal/status/pb/v1"

	"github.com/mikedonnici/mono/internal/attribute"
	attributepbv1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"
	"github.com/mikedonnici/mono/internal/status"
	"github.com/mikedonnici/mono/pkg/datastore"
)

// store keys identify specific database connections in the store
const (
	mongo1 = "mongo-1"
	mysql1 = "mysql-1"
)

func main() {

	var cfg config
	cfgFile := flag.String("cfg", "", "(optional) path to config file")
	flag.Parse()
	xf := []string{*cfgFile}
	if err := cfg.Set("mono", xf...); err != nil {
		log.Fatal(err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Set up all the store managers - in this case we're just using a single store for each manager
	// but a separate store, each with its own set of database connections, could be supplied to each manager.
	store := datastore.New()
	if err := store.AddMongoConnection(mongo1, cfg.mongoDSN, "test"); err != nil {
		return fmt.Errorf("failed to add connection for key %s: %w", mongo1, err)
	}

	if err := store.AddMySQLConnection(mysql1, cfg.mysqlDSN); err != nil {
		return fmt.Errorf("failed to add connection for key %s: %w", mysql1, err)
	}

	// Server with all the managers attached
	srvr := server{
		attributeManager: attribute.NewManager(*store),
		statusManager:    status.NewManager(*store),
	}

	s := grpc.NewServer()

	// Register the GRPC services
	attributepbv1.RegisterAttributeServiceServer(s, &srvr)
	statuspbv1.RegisterStatusServiceServer(s, &srvr)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	sn := "grpc-srvr"
	if cfg.serviceName != "" {
		sn = cfg.serviceName
	}
	log.Printf("%s listening at %v", sn, lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("could not listen: %w", err)
	}
	return nil
}
