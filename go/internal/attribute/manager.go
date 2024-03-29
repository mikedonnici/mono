package attribute

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mikedonnici/mono/pkg/datastore"
)

// Manager implements the DataManager interface
type Manager struct {

	// store provides connections to the data sources relevant to the management of attributes.
	// It is not exported and is only used by the data manager functions.
	store datastore.Connections

	// A list of data values to perform an action on
	in []Attribute

	// The resulting list once that action has been performed
	out []Attribute

	// A message relating to the last action
	msg string
}

// NewManager returns a new data manager configured thus
func NewManager(store datastore.Connections) Manager {
	return Manager{
		store: store,
	}
}

// AttributeByID fetches an attribute by id.
func (m Manager) AttributeByID(ctx context.Context, id int64) (*Attribute, error) {

	db, err := m.store.OnlyMySQLConnection()
	if err != nil {
		return nil, fmt.Errorf("conn err = %w", err)
	}

	var d Attribute
	q := "SELECT id, type, name FROM attribute WHERE id = ?"
	err = db.QueryRowContext(ctx, q, id).Scan(&d.ID, &d.Type, &d.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("no attribute with id %dn", id)
	case err != nil:
		return nil, fmt.Errorf("query error: %v", err)
	default:
		log.Printf("found attribute with id %d", id)
	}
	return &d, nil
}

func (m Manager) Create() {

}

func (m Manager) Update() {

}

//func (m Manager) Read() {
//
//	// fake fetch - this would really go to the data store
//	d := Data{
//		ID:   123,
//		Type: "ABC-123",
//		Name: "The ABC Type",
//	}
//
//	// Put in the out box
//	m.out = []Data{d}
//}
//
//func (m Manager) Delete() {
//
//}
//
//func (m Manager) Flush() {
//	m.in = []Data{}
//	m.out = []Data{}
//}
