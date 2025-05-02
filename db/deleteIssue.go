// deleteIssue.go
// Deletes one issue from the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Deletes one issue from the database with the given ID.
func (db *Database) DeleteIssue(ID string) {
	log.SetPrefix("[Delete Issue]:")

	// Get database collection.
	collection := db.Client.Database(cfg.DB_NAME).Collection(cfg.ISSUE_COLLECTION)
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println("Could not get ObjectID from input: ", err.Error())
		return
	}

	// Query filter.
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Query response.
	res := collection.FindOne(context.TODO(), filter)

	// Issue model.
	var issue *model.ProjectIssue

	// Decode response.
	_ = res.Decode(&issue)

	// Decrease open issue count if the founded issue was open.
	if issue.Status {
		// Project query filter.
		projectFilter := bson.M{"name.original": bson.M{"$eq": issue.Project}}

		// Project query response.
		projectResponse := projectCollection.FindOne(context.TODO(), projectFilter)

		// Project model.
		var project *model.Project

		// Decode response.
		err = projectResponse.Decode(&project)
		if err != nil {
			log.Println("Could not decode project: ", err.Error())
		}

		// Decrease issue count.
		reg := make(map[string]interface{})
		reg["issuecount"] = project.IssueCount - 1

		updateString := bson.M{
			"$set": reg,
		}

		// Update database project model.
		_, err = projectCollection.UpdateOne(context.TODO(), projectFilter, updateString)
		if err != nil {
			log.Println("Could not update project issue count: ", err.Error())
		}
	}

	// Delete issue.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Could not delete the issue document: ", err.Error())
		return
	}
}
