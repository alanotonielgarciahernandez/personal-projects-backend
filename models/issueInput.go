// issueInput.go
// Issue model.

package models

import "time"

// Model to create Project Issues.
type IssueInput struct {
	Date     time.Time `json:"date"`
	Labels   []*string `json:"labels"`
	Number   int       `json:"number"`
	Project  string    `json:"project"`
	RichText string    `json:"richText"`
	Status   bool      `json:"status"`
	Title    string    `json:"title"`
}
