package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// config contains all required configuration values
type config struct {

	// ServiceName is for no real rerason
	ServiceName string

	// MySQLDSN in the form: "user:pass@tcp(host:port)/dbname"
	MySQLDSN string

	errors []error
}

// Set populates the config value with vars located in the env, or in the
// specified env file(s). Env takes precedence.
func (c *config) Set(envFiles ...string) error {

	viper.SetEnvPrefix("mono")
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
	// An env var will take precedence over values from an env file.
	c.ServiceName = viper.GetString("SERVICE_NAME")
	c.MySQLDSN = viper.GetString("MYSQL_DSN")

	// Validate
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
	if c.ServiceName == "" {
		c.errors = append(c.errors, fmt.Errorf("mono_SERVICE_NAME is not set"))
	}
	if c.MySQLDSN == "" {
		c.errors = append(c.errors, fmt.Errorf("mono_MYSQL_DSN is not set"))
	}
}
