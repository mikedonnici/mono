package attribute

import pbv1 "github.com/mikedonnici/mono/internal/attribute/pb/v1"

// Attribute represents an attribute "internally", ie, in a format that is not
// used for proto messages
type Attribute struct {
	ID   int64
	Type string
	Name string
}

// Message maps an Attribute to an AttributeRaw (protobuf) message
func (a Attribute) Message() pbv1.AttributeMessage {
	return pbv1.AttributeMessage{
		Id:   a.ID,
		Type: a.Type,
		Name: a.Name,
	}
}
