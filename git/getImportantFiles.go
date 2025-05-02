// getImportantFiles.go
// Get the selected project important files list.

package git

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"errors"
	"strings"

	git "github.com/libgit2/git2go/v34"
)

// Gets the selected project important files list.
func GetImportantFiles(project *git.Repository, reference string) (map[string][]string, error) {
	// Response model.
	res := make(map[string][]string)

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
		if s != "" {
			return nil
		}

		// Iterate over important file name list to find a match.
		for _, f := range cfg.IMPORTANT_FILES {
			if f == strings.ToLower(entry.Name) {

				blob, err := project.LookupBlob(entry.Id)
				if err != nil {
					return nil
				}

				content := blob.Contents()

				contentList := strings.Split(string(content[:]), "\n")

				res[entry.Name] = contentList
			}
		}

		return nil
	})

	return res, nil
}
