package attribute

import attributev1 "github.com/mikedonnici/mono/gen/attribute/v1"

// Attribute represents an attribute "internally", ie, in a format that is not
// used for proto messages
type Attribute struct {
	ID   int64
	Type string
	Name string
}

// Message maps an Attribute to an AttributeRaw (protobuf) message
func (a Attribute) Message() attributev1.AttributeMessage {
	return attributev1.AttributeMessage{
		Id:   a.ID,
		Type: a.Type,
		Name: a.Name,
	}
}
