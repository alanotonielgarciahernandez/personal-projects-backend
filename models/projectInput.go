// projectInput.go
// Project model.

package models

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"time"
)

// Model to create Project.
type ProjectInput struct {
	Name          *model.ProjectName   `json:"name"`
	Description   string               `json:"description"`
	DefaultBranch ReferenceInput       `json:"defaultBranch"`
	BranchCount   int                  `json:"branchCount"`
	CommitCount   int                  `json:"commitCount"`
	TagCount      int                  `json:"tagCount"`
	IssueCount    int                  `json:"issueCount"`
	Images        *model.ProjectImages `json:"images"`
	Language      string               `json:"language"`
	LastUpdated   time.Time            `json:"lastUpdated"`
	Technologies  []string             `json:"technologies"`
	URL           string               `json:"url"`
}
