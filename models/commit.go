// commit.go
// Custom commit model.

package models

type LinesChanged struct {
	Added   int `json:"added"`
	Deleted int `json:"deleted"`
}

// Handler Commit response.
type Commit struct {
	Branch       string              `json:"branch"`
	Date         string              `json:"date"`
	Diff         map[string][]string `json:"diff"`
	LinesChanged LinesChanged        `json:"linesChanged"`
	Message      []string            `json:"message"`
	Parent       string              `json:"parent"`
}
