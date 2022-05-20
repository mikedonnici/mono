package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestConfig_Set(t *testing.T) {

	const envPrefix = "MONOTEST"
	const cfgFile = "./testdata/sample.cfg"
	const envServiceName = "value from env"

	// Set an env var for service name to ensure it overrides the value from the config file
	if err := os.Setenv(fmt.Sprintf("%s_SERVICE_NAME", envPrefix), envServiceName); err != nil {
		t.Fatalf("could not set env var, err = %s", err)
	}

	var cfg config
	if err := cfg.Set(envPrefix, cfgFile); err != nil {
		t.Fatalf(".Set(%s, %s) err = %s", envPrefix, cfgFile, err)
	}
	want := config{
		grpcPort:    50051,
		mongoDSN:    "mongodb://localhost:27018",
		mysqlDSN:    "root:pass@tcp(localhost:3307)/test",
		postgresDSN: "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable",
		redisDSN:    "redis://:@localhost:6380/0",
		serviceName: envServiceName,
	}
	got := cfg
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("cfg = %v, want %v", got, want)
	}
}
