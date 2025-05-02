// databaseModel.go
// Database Methods for GraphQL Server.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// ProjectRepository, Database Methods.
type ProjectRepository interface {
	DeleteIssue(ID string)
	DeleteProject(ID string)
	DeleteRelease(ID string)
	GetIssues(filter *model.IssuesFilterInput, page int) []*model.ProjectIssue
	GetProjects(fields []string, filter *model.ProjectsFilterInput, page int) []*model.Project
	GetReferences(page int, project string, referenceType string) []*model.ProjectReference
	GetRelease(project string, reference string) *model.ProjectReference
	SaveIssue(issue *model.NewIssue) string
	SaveProject(project *model.NewProject) string
	SaveRelease(release *model.NewRelease) string
	UpdateIssue(issue *model.UpdateIssue)
	UpdateProject(project *model.UpdateProject)
	UpdateRelease(release *model.UpdateRelease)
}

// Database, Contains Database Client.
type Database struct {
	Client *mongo.Client
}
