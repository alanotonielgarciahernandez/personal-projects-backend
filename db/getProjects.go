// getProjects.go
// Gets all the projects.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"log"
	"os"
	"time"

	"context"

	git "github.com/libgit2/git2go/v34"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Finds a project last update. Returns last update time.
func findLastUpdate(project string) time.Time {
	REPOSITORIES_DIR := os.Getenv("REPOSITORIES_DIR")

	repo, err := git.OpenRepository(REPOSITORIES_DIR + project + ".git")
	if err != nil {
		log.Println("Could not get repository: ", err.Error())
	}

	head, err := repo.Head()
	if err != nil {
		log.Println("Could not get repository head object: ", err.Error())
	}

	commit, err := repo.LookupCommit(head.Target())
	if err != nil {
		log.Println("Could not get repository commit object: ", err.Error())
	}

	return commit.Author().When
}

// Gets all the projects from the database. Returns projects data.
func (db *Database) GetProjects(fields []string, f *model.ProjectsFilterInput, page int) []*model.Project {
	log.SetPrefix("[Get Projects]:")

	// Response model.
	var res []*model.Project

	// Database collection.
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)

	// Query filter.
	filter := bson.M{}
	if f != nil {
		if f.Keyword != nil && len(*f.Keyword) > 0 {
			filter["name.original"] = bson.M{"$regex": *f.Keyword, "$options": "i"}
		}
	}

	// Query options.
	options := options.FindOptions{}
	options.SetLimit(10)
	options.SetSkip(10 * (int64(page) - 1))

	// Query response.
	projectCursor, err := projectCollection.Find(context.TODO(), filter, &options)
	if err != nil {
		log.Println("Could not find project documents: ", err.Error())
	}
	defer projectCursor.Close(context.TODO())

	// Build the project list.
	for projectCursor.Next(context.TODO()) {
		// Project model.
		var project *model.Project

		// Decode response.
		err := projectCursor.Decode(&project)
		if err != nil {
			log.Println("Could not decode project: ", err.Error())
		}

		// Find project last time updated if requested.
		for i := range fields {
			if fields[i] == "lastUpdated" {
				project.LastUpdated = findLastUpdate(project.Name.Original)
			}
		}

		res = append(res, project)
	}

	return res
}
