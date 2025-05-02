// updateProject.go
// Updates one project.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Updates one project from the database with the given ID.
func (db *Database) UpdateProject(p *model.UpdateProject) {
	log.SetPrefix("[Update Project]:")

	// Database collection.
	collection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)

	// Build update keys.
	reg := make(map[string]interface{})
	if p.Description != nil {
		reg["description"] = p.Description
	}
	if p.Images != nil {
		reg["images"] = p.Images
	}
	if p.Language != nil {
		reg["language"] = p.Language
	}
	if p.Name != nil {
		reg["name"] = p.Name
	}
	if p.URL != nil {
		reg["url"] = p.URL
	}

	// Update model.
	updString := bson.M{
		"$set": reg,
	}

	objID, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		log.Println("Could not get ObjectID from input: ", err.Error())
		return
	}

	// Query filter.
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Update project.
	_, err = collection.UpdateOne(context.TODO(), filter, updString)
	if err != nil {
		log.Println("Could not update the project document: ", err.Error())
		return
	}
}
