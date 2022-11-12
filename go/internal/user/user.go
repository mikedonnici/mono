package user

import pbv1 "github.com/mikedonnici/mono/internal/user/pb/v1"

// User represents a user inside the API boundary
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
}

// Message maps the internal User to the protobuf UserMessage
func (u User) Message() pbv1.UserMessage {
	return pbv1.UserMessage{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
