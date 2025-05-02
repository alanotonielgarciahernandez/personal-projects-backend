// deleteRelease.go
// Deletes one release from the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Deletes one release model from the database with the given ID.
func (db *Database) DeleteRelease(ID string) {
	log.SetPrefix("[Delete Release]:")

	// Get database collection.
	collection := db.Client.Database(cfg.DB_NAME).Collection(cfg.RELEASE_COLLECTION)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println("Could not get ObjectID from input: ", err.Error())
		return
	}

	// Query filter.
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Delete release.
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Could not delete the release document: ", err.Error())
		return
	}
}
