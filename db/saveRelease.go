// saveRelease.go
// Saves one release into the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Saves one release into the database with the given data.
// Returns the inserted release ID.
func (db *Database) SaveRelease(release *model.NewRelease) string {
	log.SetPrefix("[Save Release]:")

	// Database collection.
	collection := db.Client.Database(cfg.DB_NAME).Collection(cfg.RELEASE_COLLECTION)

	tagID, err := primitive.ObjectIDFromHex(release.Tag)
	if err != nil {
		log.Println("Could not get ObjectID from input: ", err.Error())
	}

	// Release model.
	newRelease := models.ReleaseInput{
		Project:  release.Project,
		RichText: release.RichText,
		Tag:      tagID,
	}

	// Insert model into database.
	res, err := collection.InsertOne(context.TODO(), newRelease)
	if err != nil {
		log.Println("Could not insert release: ", err.Error())
	}

	return res.InsertedID.(primitive.ObjectID).Hex()
}
