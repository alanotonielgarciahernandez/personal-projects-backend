// updateIssue.go
// Updates one issue.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Updates one issue from the database with the given ID.
func (db *Database) UpdateIssue(i *model.UpdateIssue) {
	log.SetPrefix("[Update Issue]:")

	// Database collections.
	issueCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.ISSUE_COLLECTION)
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)

	issueID, err := primitive.ObjectIDFromHex(i.ID)
	if err != nil {
		log.Println("Could not get ObjectID from ID field: ", err.Error())
	}

	// Query filter.
	filter := bson.M{"_id": bson.M{"$eq": issueID}}

	// Build update keys.
	reg := make(map[string]interface{})
	if i.Labels != nil {
		reg["labels"] = i.Labels
	}
	if i.Project != nil {
		projectID, err := primitive.ObjectIDFromHex(*i.Project)
		if err != nil {
			log.Println("Could not get ObjectID from Project field: ", err.Error())
		}
		reg["project"] = projectID
	}
	if i.RichText != nil {
		reg["richtext"] = i.RichText
	}
	if i.Title != nil {
		reg["title"] = i.Title
	}
	if i.Number != nil {
		reg["number"] = i.Number
	}
	if i.Status != nil {
		reg["status"] = i.Status

		status := *i.Status
		var issue *model.ProjectIssue

		findRes := issueCollection.FindOne(context.TODO(), filter)
		findRes.Decode(&issue)

		projectFilter := bson.M{"name.original": bson.M{"$eq": issue.Project}}

		projectResponse := projectCollection.FindOne(context.TODO(), projectFilter)
		var project *model.Project

		err = projectResponse.Decode(&project)
		if err != nil {
			log.Println("Could not decode project: ", err.Error())
		}
		registry := make(map[string]interface{})

		// Update open issue count.
		if issue.Status && !status {
			registry["issuecount"] = project.IssueCount - 1
		} else if !issue.Status && status {
			registry["issuecount"] = project.IssueCount + 1
		}

		updateString := bson.M{
			"$set": registry,
		}

		_, err = projectCollection.UpdateOne(context.TODO(), projectFilter, updateString)
		if err != nil {
			log.Println("Could not update project issue count: ", err.Error())
		}
	}

	// Update model.
	updString := bson.M{
		"$set": reg,
	}

	// Update issue.
	_, err = issueCollection.UpdateOne(context.TODO(), filter, updString)
	if err != nil {
		log.Println("Could not update the issue document: ", err.Error())
		return
	}
}
