// getRelease.go
// Gets a project release data.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Gets a release object from database.
func (db *Database) GetRelease(project string, reference string) *model.ProjectReference {
	log.SetPrefix("[Get Release]:")

	// Database collections.
	referenceCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.REFERENCE_COLLECTION)
	releaseCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.RELEASE_COLLECTION)

	// Query filter.
	filter := bson.M{
		"project": bson.M{"$eq": project},
		"name":    bson.M{"$eq": reference},
	}

	// Query response.
	releaseResponse := releaseCollection.FindOne(context.TODO(), filter)

	// Release model.
	var release *model.ProjectReference

	// Decode response.
	err := releaseResponse.Decode(&release)
	if err != nil {
		referenceResponse := referenceCollection.FindOne(context.TODO(), filter)
		err := referenceResponse.Decode(&release)
		if err != nil {
			log.Println("Could not decode reference: ", err.Error())
		}
	}

	return release
}
