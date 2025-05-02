// updateRelease.go
// Updates one release.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Updates one project release from the database with the given ID.
func (db *Database) UpdateRelease(r *model.UpdateRelease) {
	log.SetPrefix("[Update Release]:")

	// Database collection.
	collection := db.Client.Database(cfg.DB_NAME).Collection(cfg.RELEASE_COLLECTION)

	releaseID, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		log.Println("Could not get ObjectID from ID field: ", err.Error())
	}

	// Build update keys.
	reg := make(map[string]interface{})
	if r.Project != nil {
		projectID, err := primitive.ObjectIDFromHex(*r.Project)
		if err != nil {
			log.Println("Could not get ObjectID from Project field: ", err.Error())
		}
		reg["project"] = projectID
	}
	if r.RichText != nil {
		reg["richtext"] = r.RichText
	}
	if r.Tag != nil {
		reg["tag"] = r.Tag
	}

	// Update model.
	updString := bson.M{
		"$set": reg,
	}

	// Query filter.
	filter := bson.M{"_id": bson.M{"$eq": releaseID}}

	// Update release.
	_, err = collection.UpdateOne(context.TODO(), filter, updString)
	if err != nil {
		log.Println("Could not update the release document: ", err.Error())
		return
	}
}
