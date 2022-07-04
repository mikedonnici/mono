package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Sensible defaults
const (
	defaultGRPCPort = 50051
)

// config contains all required configuration values
type config struct {
	serviceName string // no real purpose
	grpcPort    int
	mongoDSN    string // mongodb://host:port
	mysqlDSN    string // user:pass@tcp(host:port)/dbname
	postgresDSN string // postgres://user:pass@host:port/dbname?sslmode=disable
	redisDSN    string // redis://user:pass@host:port/[dbnumber]
	errors      []error
}

// Set populates the config value with vars located in the env, or in the specified env file(s) - env takes precedence
// Override prefix for tests so real env does not create issues.
func (c *config) Set(prefix string, envFiles ...string) error {

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	for _, envFile := range envFiles {
		viper.SetConfigName(envFile)
		viper.SetConfigType("env")
	}
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("could not read config file(s) - hopefully values exists in the environment!")
	}

	// These .Get calls trigger .AutomaticEnv() which should fetch a value from the environment, if it exists.
	// An env var has precedence over values from env/cfg files and env vars are expected to have the prefix.
	c.grpcPort = viper.GetInt("GRPC_PORT")
	c.mongoDSN = viper.GetString("MONGO_DSN")
	c.mysqlDSN = viper.GetString("MYSQL_DSN")
	c.postgresDSN = viper.GetString("POSTGRES_DSN")
	c.redisDSN = viper.GetString("REDIS_DSN")
	c.serviceName = viper.GetString("SERVICE_NAME")

	c.validate()
	if len(c.errors) > 0 {
		var msg string
		for _, err := range c.errors {
			msg += fmt.Sprintf("%s%s", err, "\n")
		}
		return fmt.Errorf("config.errors has %d errors:\n%s", len(c.errors), msg)
	}
	return nil
}

func (c *config) validate() {
	if c.grpcPort == 0 {
		c.grpcPort = defaultGRPCPort
		log.Printf("grpc port set to default %d", defaultGRPCPort)
	}
	if c.mongoDSN == "" {
		c.errors = append(c.errors, fmt.Errorf("MONGO_DSN is not set"))
	}
	if c.mysqlDSN == "" {
		c.errors = append(c.errors, fmt.Errorf("MYSQL_DSN is not set"))
	}
	if c.postgresDSN == "" {
		c.errors = append(c.errors, fmt.Errorf("POSTGRES_DSN is not set"))
	}
	if c.redisDSN == "" {
		c.errors = append(c.errors, fmt.Errorf("REDIS_DSN is not set"))
	}
}
