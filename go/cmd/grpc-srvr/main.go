package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mikedonnici/mono/internal/attribute"
	attributepbv1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"
	"github.com/mikedonnici/mono/internal/status"
	statuspbv1 "github.com/mikedonnici/mono/internal/status/pb/v1"
	"github.com/mikedonnici/mono/internal/user"
	userpbv1 "github.com/mikedonnici/mono/internal/user/pb/v1"
	"github.com/mikedonnici/mono/pkg/datastore"
)

// store keys identify specific database connections in the store
const (
	mongo1 = "mongo1"
	mysql1 = "mysql1"
	redis1 = "redis1"
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
	// Set up the required data connections. Using a single store here but could use separate stores for each
	// data manager to ensure isolation between components.
	store, err := dataConnections(cfg)
	if err != nil {
		return fmt.Errorf("failed to create data store: %w", err)
	}

	// Attach data managers to a server
	srvr := server{
		attributeManager: attribute.NewManager(*store),
		statusManager:    status.NewManager(*store),
		userManager:      user.NewManager(*store),
	}

	// Register GRPC services with the server
	s := grpc.NewServer()
	attributepbv1.RegisterAttributeServiceServer(s, &srvr)
	statuspbv1.RegisterStatusServiceServer(s, &srvr)
	userpbv1.RegisterUserServiceServer(s, &srvr)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Set up listener and start listening!
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.grpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Printf("listening at p %v", lis.Addr())
	return nil
}

func dataConnections(cfg config) (*datastore.Connections, error) {
	store := datastore.New()
	if err := store.AddRedisConnection(redis1, cfg.redisDSN, 30); err != nil {
		return nil, fmt.Errorf("failed to add redis connection with key %s: %w", redis1, err)
	}
	//if err := store.AddMongoConnection(mongo1, cfg.mongoDSN, "test"); err != nil {
	//	return fmt.Errorf("failed to add mongodb connection with key %s: %w", mongo1, err)
	//}
	//
	if err := store.AddMySQLConnection(mysql1, cfg.mysqlDSN); err != nil {
		return nil, fmt.Errorf("failed to add mysql connection with key %s: %w", mysql1, err)
	}
	return store, nil
}
