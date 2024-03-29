package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	DSN string
	*sql.DB
}

// newPostgresConnection returns a pointer to a PostgresConnection with an established
// db session, or an error.
func newPostgresConnection(dsn string) (*PostgresConnection, error) {
	conn := PostgresConnection{
		DSN: dsn,
		DB:  nil,
	}
	db, err := conn.open()
	if err != nil {
		return nil, fmt.Errorf("could not open postgres db connection, err = %w", err)
	}
	conn.DB = db
	if err := conn.Check(); err != nil {
		return nil, fmt.Errorf("postgres db connection check failed, err = %w", err)
	}
	return &conn, nil
}

// connectPostgres returns a postgres connection or an error.
func connectPostgres(dsn string) (*PostgresConnection, error) {
	return newPostgresConnection(dsn)
}

// open returns a connection to the database or an error.
func (c *PostgresConnection) open() (*sql.DB, error) {
	return sql.Open("postgres", c.DSN)
}

// Check verifies the connection to the database and returns an error if there's a problem.
// Note: This is better than ping because it forces a round trip to the database.
func (c *PostgresConnection) Check() error {
	var tmp bool
	return c.DB.QueryRow(`SELECT true`).Scan(&tmp)
}

// connectRedis returns a redis connection or an error.
func connectRedis(dsn string, timeoutSeconds int) (*RedisConnection, error) {
	return newRedisConnection(dsn, timeoutSeconds)
}

// // open returns a connection to the database or an error.
// func (c *postgresConnection) open() (*sql.DB, error) {
//	return sql.Open("postgres", c.DSN)
// }
//
// // Check verifies the connection to the database and returns an error if there's a problem.
// // Note: This is better than ping because it forces a round trip to the database.
// func (c *postgresConnection) Check() error {
//	var tmp bool
//	return c.DB.QueryRow(`SELECT true`).Scan(&tmp)
// }
