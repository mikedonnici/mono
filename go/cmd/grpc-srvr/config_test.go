package main

import (
	"reflect"
	"testing"
)

func TestConfig_Set(t *testing.T) {

	const envPrefix = "MONOTEST"
	const cfgFile = "./testdata/sample.cfg"

	var cfg config
	if err := cfg.Set(envPrefix, cfgFile); err != nil {
		t.Fatalf(".Set(%s, %s) err = %s", envPrefix, cfgFile, err)
	}
	want := config{
		grpcPort: 50051,
		//mongoDSN:    "mongodb://localhost:27018",
		mysqlDSN: "root:pass@tcp(localhost:3306)/test",
		//postgresDSN: "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable",
		redisDSN: "redis://:@localhost:6380/0",
	}
	got := cfg
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("cfg = %v, want %v", got, want)
	}
}
