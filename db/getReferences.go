// getReferences.go
// Gets all the project references.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Gets all the references from database projects. Returns a reference list.
func (db *Database) GetReferences(page int, project string, referenceType string) []*model.ProjectReference {
	log.SetPrefix("[Get References]:")

	// Response model.
	var res = []*model.ProjectReference{}

	// Database collection.
	referenceCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.REFERENCE_COLLECTION)

	// Query filter.
	filter := bson.M{
		"project": bson.M{"$eq": project},
		"type":    bson.M{"$eq": referenceType},
	}

	// Query options.
	options := &options.FindOptions{}
	if page > 0 {
		options.SetLimit(10)
		options.SetSkip(10 * (int64(page) - 1))
	}
	options.SetSort(bson.M{"date": -1})

	// Query response.
	referenceCursor, err := referenceCollection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Println("Could not find reference documents: ", err.Error())
	}
	defer referenceCursor.Close(context.TODO())

	// Build the reference list.
	for referenceCursor.Next(context.TODO()) {
		// References model.
		var reference *model.ProjectReference

		// Decode response.
		err := referenceCursor.Decode(&reference)
		if err != nil {
			log.Println("Could not decode reference: ", err.Error())
		}

		res = append(res, reference)
	}

	return res
}
