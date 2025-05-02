// releaseInput.go
// Project release model.

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model to create Project Releases.
type ReleaseInput struct {
	Project  string             `json:"project"`
	RichText string             `json:"richText"`
	Tag      primitive.ObjectID `json:"tag"`
}
