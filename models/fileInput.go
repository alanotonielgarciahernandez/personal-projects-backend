// fileInput.go
// Repository file model.

package models

// Model to create repository files.
type FileInput struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Path        string               `json:"path"`
	Type        string               `json:"type"`
	Tree        map[string]FileInput `json:"tree"`
}
