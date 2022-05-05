// Package datastore provides a structure within which to store connections to multiple data sources and services.
package datastore

import (
	"fmt"
	"log"
)

// Connections represents connections to various platform services, most likely databases.
type Connections struct {
	Mongo    map[string]*mongoConnection
	MySQL    map[string]*mysqlConnection
	Postgres map[string]*postgresConnection
	Redis    map[string]*redisConnection
}

// New returns a pointer to a Connections value no connections attached.
func New() *Connections {
	return &Connections{
		Mongo:    make(map[string]*mongoConnection),
		MySQL:    make(map[string]*mysqlConnection),
		Postgres: make(map[string]*postgresConnection),
		Redis:    make(map[string]*redisConnection),
	}
}

// AddMongoConnection adds a connection to a mongo database identified by the specified key.
func (mc *Connections) AddMongoConnection(key, dsn, db string) error {
	log.Printf("Adding MONGO connection: key = '%s', db = '%s'", key, db)
	c, err := connectMongo(dsn, db)
	if err != nil {
		return fmt.Errorf("could not connect to mongo, err = %w", err)
	}
	mc.Mongo[key] = c
	return nil
}

// MongoConnByKey returns the mongodb connection value at the specified key.
func (mc *Connections) MongoConnByKey(key string) (*mongoConnection, error) {
	c, ok := mc.Mongo[key]
	if !ok {
		return nil, fmt.Errorf("no mongodb connection with key = %s", key)
	}
	return c, nil
}

// OnlyMongoConnection is a convenience function that returns the Mongo connection if there is only one.
func (mc *Connections) OnlyMongoConnection() (*mongoConnection, error) {
	num := len(mc.Mongo)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique mongo connection as %d exist", num)
	}
	var key string
	for k := range mc.Mongo { // there's only one
		key = k
	}
	return mc.Mongo[key], nil
}

// AddMySQLConnection adds a connection to a mysql database identified by the specified key.
func (mc *Connections) AddMySQLConnection(key, dsn string) error {
	log.Printf("Adding MYSQL connection: key = '%s'", key)
	c, err := connectMySQL(dsn)
	if err != nil {
		return fmt.Errorf("could not connect to MySQL, err = %w", err)
	}
	mc.MySQL[key] = c
	return nil
}

// MySQLConnByKey returns the MySQL connection value at the specified key.
func (mc *Connections) MySQLConnByKey(key string) (*mysqlConnection, error) {
	c, ok := mc.MySQL[key]
	if !ok {
		return nil, fmt.Errorf("no MySQL db connection with key = %s", key)
	}
	return c, nil
}

// OnlyMySQLConnection is a convenience function that returns the MySQL connection if there is only one.
func (mc *Connections) OnlyMySQLConnection() (*mysqlConnection, error) {
	num := len(mc.MySQL)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique MySQL connection as %d exist", num)
	}
	var key string
	for k := range mc.MySQL {
		key = k
	}
	return mc.MySQL[key], nil
}

// AddPostgresConnection adds a connection to a potsgres database identified by the specified key.
func (mc *Connections) AddPostgresConnection(key, dsn string) error {
	log.Printf("Adding POSTRES connection: key = '%s'", key)
	c, err := connectPostgres(dsn)
	if err != nil {
		return fmt.Errorf("could not connect to postgres, err = %w", err)
	}
	mc.Postgres[key] = c
	return nil
}

// PostgresConnByKey returns the postgres connection value at the specified key.
func (mc *Connections) PostgresConnByKey(key string) (*postgresConnection, error) {
	c, ok := mc.Postgres[key]
	if !ok {
		return nil, fmt.Errorf("no postgresdb connection with key = %s", key)
	}
	return c, nil
}

// OnlyPostgresConnection is a convenience function that returns the Postgres connection if there is only one.
func (mc *Connections) OnlyPostgresConnection() (*postgresConnection, error) {
	num := len(mc.Postgres)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique postgres connection as %d exist", num)
	}
	var key string
	for k := range mc.Postgres {
		key = k
	}
	return mc.Postgres[key], nil
}

// AddRedisConnection adds a connection to a redis database identified by the specified key.
func (mc *Connections) AddRedisConnection(key, dsn string, timeoutSeconds int) error {
	log.Printf("Adding REDIS connection: key = '%s'", key)
	c, err := connectRedis(dsn, timeoutSeconds)
	if err != nil {
		return fmt.Errorf("could not connect to Redis, err = %w", err)
	}
	mc.Redis[key] = c
	return nil
}

// RedisConnByKey returns the Redis connection value at the specified key.
func (mc *Connections) RedisConnByKey(key string) (*redisConnection, error) {
	c, ok := mc.Redis[key]
	if !ok {
		return nil, fmt.Errorf("no Redisdb connection with key = %s", key)
	}
	return c, nil
}

// OnlyRedisConnection is a convenience function that returns the Redis connection if there is only one.
func (mc *Connections) OnlyRedisConnection() (*redisConnection, error) {
	num := len(mc.Redis)
	if num != 1 {
		return nil, fmt.Errorf("cannot return unique Redis connection as %d exist", num)
	}
	var key string
	for k := range mc.Redis {
		key = k
	}
	return mc.Redis[key], nil
}
