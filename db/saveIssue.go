// saveIssue.go
// Saves an issue into the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Saves one issue to the database with the given data.
// Returns the inserted issue ID.
func (db *Database) SaveIssue(issue *model.NewIssue) string {
	log.SetPrefix("[Save Issue]:")

	// Database collections.
	issueCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.ISSUE_COLLECTION)
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)

	// Query filters.
	projectFilter := bson.M{"name.original": bson.M{"$eq": issue.Project}}
	countFilter := bson.M{"project": bson.M{"$eq": issue.Project}}

	// Count issues.
	count, err := issueCollection.CountDocuments(context.TODO(), countFilter, options.Count().SetHint("_id_"))
	if err != nil {
		log.Println("Could not count issue documents: ", err.Error())
	}

	// Issue model.
	newIssue := models.IssueInput{
		Date:     time.Now(),
		Labels:   issue.Labels,
		Number:   int(count + 1),
		Project:  issue.Project,
		RichText: issue.RichText,
		Status:   issue.Status,
		Title:    issue.Title,
	}

	// Insert issue.
	res, err := issueCollection.InsertOne(context.TODO(), newIssue)
	if err != nil {
		log.Println("Could not insert issue: ", err.Error())
	}

	// If the inserted issue is created as open, increase the open issue count.
	if issue.Status {
		// Query response.
		projectResponse := projectCollection.FindOne(context.TODO(), projectFilter)

		// Project model.
		var project *model.Project

		// Decode response.
		err = projectResponse.Decode(&project)
		if err != nil {
			log.Println("Could not decode project: ", err.Error())
		}

		// Increase the open issue count.
		reg := make(map[string]interface{})
		reg["issuecount"] = project.IssueCount + 1

		updateString := bson.M{
			"$set": reg,
		}

		// Update the database issue model.
		_, err = projectCollection.UpdateOne(context.TODO(), projectFilter, updateString)
		if err != nil {
			log.Println("Could not update project issue count: ", err.Error())
		}
	}

	return res.InsertedID.(primitive.ObjectID).Hex()
}
