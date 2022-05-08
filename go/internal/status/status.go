package status

import "github.com/mikedonnici/mono/pkg/datastore"

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
