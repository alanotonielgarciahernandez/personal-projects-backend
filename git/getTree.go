// getTree.go
// Gets project files tree.

package git

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"errors"
	"strings"

	git "github.com/libgit2/git2go/v34"
)

// Gets the selected project files tree at the selected reference.
// Returns a map of the tree.
func GetTree(project *git.Repository, reference string, path string) (map[string]models.FileInput, error) {
	// Response model.
	res := make(map[string]models.FileInput)

	oid, err := git.NewOid(reference)
	if err != nil {
		return res, errors.New("reference is not valid, " + err.Error())
	}

	commit, err := project.LookupCommit(oid)
	if err != nil {
		return res, errors.New("while getting reference, " + err.Error())
	}

	currentTree, err := commit.Tree()
	if err != nil {
		return res, errors.New("while getting tree, " + err.Error())
	}

	// Walk tree.
	currentTree.Walk(func(s string, entry *git.TreeEntry) error {
		// File model.
		file := models.FileInput{
			Name:        entry.Name,
			Description: "",
			Path:        s,
			Type:        entry.Type.String(),
		}

		// Create folder subtree.
		if entry.Type.String() == "Tree" {
			file.Tree = make(map[string]models.FileInput)
		}

		// Find file tree.
		current := &res

		if file.Path == "" {
			res[file.Name] = file
		} else {
			for _, segment := range strings.Split(s, "/") {
				for _, v := range *current {
					if v.Name == segment {
						current = &v.Tree
						break
					}
				}
			}

			(*current)[file.Name] = file
		}

		return nil
	})

	return res, nil
}
