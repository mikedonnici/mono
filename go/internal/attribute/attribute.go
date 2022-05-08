package attribute

// Attribute represents an attribute "internally", ie, in a format that is not
// used for proto messages
type Attribute struct {
	ID   int64
	Type string
	Name string
}

// raw maps an Attribute to an AttributeRaw (protobuf) message
func (a Attribute) Raw() AttributeRaw {
	return AttributeRaw{
		Id:   a.ID,
		Type: a.Type,
		Name: a.Name,
	}
}
