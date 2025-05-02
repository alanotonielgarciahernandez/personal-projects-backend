// deleteProject.go
// Deletes one project from the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Deletes one project from the database with the given ID.
func (db *Database) DeleteProject(projectName string) {
	log.SetPrefix("[Delete Project]:")

	// Get database collections.
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)
	referenceCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.REFERENCE_COLLECTION)

	// Query filters.
	idFilter := bson.M{"name.original": bson.M{"$eq": projectName}}
	projectFilter := bson.M{"project": bson.M{"$eq": projectName}}

	// Delete database project model.
	_, err := projectCollection.DeleteOne(context.TODO(), idFilter)
	if err != nil {
		log.Println("Could not delete the project document: ", err.Error())
		return
	}

	// Delete database project references.
	_, err = referenceCollection.DeleteMany(context.TODO(), projectFilter)
	if err != nil {
		log.Println("Could not delete the reference document: ", err.Error())
		return
	}
}
