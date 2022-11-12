package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mikedonnici/mono/pkg/datastore"
)

// Manager implements the DataManager interface
type Manager struct {
	store datastore.Connections
}

// NewManager returns a new data manager configured thus
func NewManager(store datastore.Connections) Manager {
	return Manager{
		store: store,
	}
}

// UserByID fetches a user by id
func (m Manager) UserByID(ctx context.Context, id int64) (*User, error) {
	db, err := m.store.OnlyMySQLConnection()
	if err != nil {
		return nil, fmt.Errorf("conn err = %w", err)
	}

	var u User
	q := "SELECT id, firstname, lastname, email FROM user WHERE id = ?"
	err = db.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("no user with id %d", id)
	case err != nil:
		return nil, fmt.Errorf("query error: %v", err)
	}
	return &u, nil
}
