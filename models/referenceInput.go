// referenceInput.go
// Project reference model.

package models

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"time"
)

// Model to create Project References.
type ReferenceInput struct {
	Hash        string             `json:"hash"`
	Project     string             `json:"project"`
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	Date        time.Time          `json:"date"`
	Default     *bool              `json:"default,omitempty"`
	Aheadbehind *model.AheadBehind `json:"aheadbehind,omitempty"`
	Message     []string           `json:"message,omitempty"`
}
