package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestConfig_Set(t *testing.T) {

	f := "temp.env"
	envPrefix := "MONOTEST"
	serviceName := "value from env"

	// Write a temp env file with DSN and with a service name
	// env file values will NOT have the prefix, but env vars WILL have the prefix
	envFile := "MYSQL_DSN=user:pass@tcp(host:port)/dbname\nSERVICE_NAME=\"value from file\""
	if err := os.WriteFile(f, []byte(envFile), 0666); err != nil {
		t.Fatalf("could not create temp env file '%s', err = %s", f, err)
	}

	// Service name can come from an env var - env vars WILL have the prefix and should
	// override the value in the env file created above
	if err := os.Setenv(fmt.Sprintf("%s_SERVICE_NAME", envPrefix), serviceName); err != nil {
		t.Fatalf("could not set env var, err = %s", err)
	}

	var cfg config
	if err := cfg.Set(envPrefix, f); err != nil {
		t.Fatalf(".Set() err = %s", err)
	}
	want := config{
		grpcPort:    defaultGRPCPort,
		serviceName: serviceName,
		mysqlDSN:    "user:pass@tcp(host:port)/dbname",
	}
	got := cfg
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("cfg = %v, want %v", got, want)
	}

	t.Log(cfg)

	if err := os.Remove(f); err != nil {
		t.Fatalf("could not remove temp env file '%s', err = %s", f, err)
	}
}
