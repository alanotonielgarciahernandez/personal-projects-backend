// saveProject.go
// Saves one project into the database.

package db

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/git"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	git2go "github.com/libgit2/git2go/v34"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Saves one project to the database with the given data.
// Returns the inserted project ID.
func (db *Database) SaveProject(project *model.NewProject) string {
	log.SetPrefix("[Save Project]:")

	// References list.
	var refCol []interface{}

	// Project model.
	newProject := models.ProjectInput{
		Name:         (*model.ProjectName)(project.Name),
		Description:  project.Description,
		IssueCount:   0,
		Language:     project.Language,
		Technologies: project.Technologies,
		URL:          project.URL,
	}

	// Database collections.
	projectCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.PROJECT_COLLECTION)
	referenceCollection := db.Client.Database(cfg.DB_NAME).Collection(cfg.REFERENCE_COLLECTION)

	REPOSITORIES_DIR := os.Getenv("REPOSITORIES_DIR")

	proj, err := git2go.OpenRepository(REPOSITORIES_DIR + project.Name.Original + ".git")
	if err != nil {
		log.Fatal("Could not open project: ", err.Error())
	}

	// Create uploads directory.
	err = os.MkdirAll(cfg.UPLOADS_FOLDER, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create uploads directory: ", err.Error())
	}

	// Create images name.
	backgroundName := project.Name.Original + "-background" + filepath.Ext(project.Images.Background.Filename)
	logoName := project.Name.Original + "-logo" + filepath.Ext(project.Images.Logo.Filename)

	// Create image files into the uploads directory.
	backgroundOut, err := os.Create(filepath.Join(cfg.UPLOADS_FOLDER, backgroundName))
	if err != nil {
		log.Fatal("Could not create background image file: ", err.Error())
	}
	defer backgroundOut.Close()

	logoOut, err := os.Create(filepath.Join(cfg.UPLOADS_FOLDER, logoName))
	if err != nil {
		log.Fatal("Could not create logo image file: ", err.Error())
	}
	defer logoOut.Close()

	// Copy provided images data into the new image files.
	if _, err := io.Copy(backgroundOut, project.Images.Background.File); err != nil {
		log.Fatal("Could not insert background image file: ", err.Error())
	}

	if _, err := io.Copy(logoOut, project.Images.Logo.File); err != nil {
		log.Fatal("Could not insert logo image file: ", err.Error())
	}

	// Set images routes to the project model.
	newProject.Images = &model.ProjectImages{
		Background: fmt.Sprintf("%s/%s", cfg.UPLOADS_FOLDER, backgroundName),
		Logo:       fmt.Sprintf("%s/%s", cfg.UPLOADS_FOLDER, logoName),
	}

	// Get references.
	projectReferences, branchCount, commitCount, tagCount, head, err := git.GetReferences(proj, project.Name.Original)
	if err != nil {
		log.Fatal("Could not get project references: ", err.Error())
	}

	// Set reference variables.
	newProject.BranchCount = branchCount
	newProject.CommitCount = commitCount
	newProject.TagCount = tagCount
	newProject.LastUpdated = projectReferences[0].Date
	newProject.DefaultBranch = head

	// Build reference list.
	for _, r := range projectReferences {
		refCol = append(refCol, r)
	}

	// Insert project references.
	_, err = referenceCollection.InsertMany(context.TODO(), refCol)
	if err != nil {
		log.Fatal("Could not insert references: ", err.Error())
	}

	// Insert project model.
	insertResult, err := projectCollection.InsertOne(context.TODO(), newProject)
	if err != nil {
		log.Fatal("Could not insert project: ", err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex()
}
