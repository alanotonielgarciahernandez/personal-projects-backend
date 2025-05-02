// getIssues.go
// Gets all the project issues.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"log"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Gets the project issues from the database. Returns issues data.
func (db *Database) GetIssues(f *model.IssuesFilterInput, page int) []*model.ProjectIssue {
	log.SetPrefix("[Get Issues]:")

	// Response model.
	var res []*model.ProjectIssue

	// Database collection.
	issueCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.ISSUE_COLLECTION)

	// Query filter.
	filter := bson.M{}
	if f != nil {
		if f.Project != nil && len(*f.Project) > 0 {
			filter["project"] = bson.M{"$eq": f.Project}
		}
		if f.Status != nil {
			filter["status"] = bson.M{"$eq": f.Status}
		}
		if f.Number != nil {
			filter["number"] = bson.M{"$eq": f.Number}
		}
	}

	// Query options.
	options := options.FindOptions{}
	options.SetLimit(10)
	options.SetSkip(10 * (int64(page) - 1))

	// Query response.
	issueCursor, err := issueCollection.Find(context.TODO(), filter, &options)
	if err != nil {
		log.Println("Could not find issue documents: ", err.Error())
	}
	defer issueCursor.Close(context.TODO())

	// Build issue list.
	for issueCursor.Next(context.TODO()) {
		// Issue model.
		var issue *model.ProjectIssue

		// Decode response.
		err := issueCursor.Decode(&issue)
		if err != nil {
			log.Println("Could not decode issue: ", err.Error())
		}

		res = append(res, issue)
	}

	return res
}
