package attribute

// Config contains values required to configure the attribute data manager
type Config struct {
	dsn string
}

// Data represents an attribute internally. This needs to be mapped
// to the Attribute message at some point.
// This is ok given that the internal format may need to be 'massaged' into its final shape.
// This also allows validation to take place at this internal data level, before we map it to a message.
// Todo: try custom tags to make mapping easier?
type Data struct {
	ID   int64
	Type string
	Name string
}

// asMessage returns Data as a proto message
func (d Data) asMessage() Attribute {
	return Attribute{
		Id:   d.ID,
		Type: d.Type,
		Name: d.Name,
	}
}

// Manager implements the DataManager interface
type Manager struct {
	cfg Config

	// This type will be an actual data store that is relevant to the
	// management of attributes. It is not exported and is only used
	// by the data manager functions.
	store string

	// A list of data values to perform an action on
	in []Data

	// The resulting list once that action has been performed
	out []Data

	// A message relating to the last action
	msg string
}

// NewManager returns a new data manager configured thus
func NewManager(cfg Config) Manager {
	m := Manager{
		cfg:   cfg,
		store: "this would be an actual connection",
	}
	return m
}

// AttributeByID fetches an attribute by id.
func (m Manager) AttributeByID(id int64) (Attribute, error) {

	d := Data{
		ID:   id,
		Type: "ABC-123",
		Name: "The ABC Type",
	}
	return d.asMessage(), nil
}

func (m Manager) Create() {

}

func (m Manager) Update() {

}

func (m Manager) Read() {

	// fake fetch - this would really go to the data store
	d := Data{
		ID:   123,
		Type: "ABC-123",
		Name: "The ABC Type",
	}

	// Put in the out box
	m.out = []Data{d}
}

func (m Manager) Delete() {

}

func (m Manager) Flush() {
	m.in = []Data{}
	m.out = []Data{}
}
