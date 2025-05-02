// getContent.go
// Gets a file content.

package git

import (
	"errors"
	"os"
	"strings"

	git "github.com/libgit2/git2go/v34"
)

// Gets the selected file content. Returns a list of the content lines.
func GetContent(project string, filePath string, reference string) ([]string, error) {
	// Variables.
	var commit *git.Commit
	var oid *git.Oid

	REPOSITORIES_DIR := os.Getenv("REPOSITORIES_DIR")

	if REPOSITORIES_DIR == "" {
		return nil, errors.New("missing environment variable REPOSITORIES_DIR")
	}

	repo, err := git.OpenRepository(REPOSITORIES_DIR + project + ".git")
	if err != nil {
		return nil, err
	}

	// Get project head references if not specified.
	if reference == "" {
		ref, err := repo.Head()
		if err != nil {
			return []string{}, err
		}

		commit, err = repo.LookupCommit(ref.Target())
		if err != nil {
			return []string{}, err
		}
	} else {
		// Find reference hash.
		oid, err = git.NewOid(reference)
		if err != nil {
			ref, err := repo.References.Lookup("refs/heads/" + reference)
			if err != nil {
				ref, err = repo.References.Lookup("refs/tags/" + reference)
				if err != nil {
					return []string{}, err
				}
			}

			oid = ref.Target()
		}

		commit, err = repo.LookupCommit(oid)
		if err != nil {
			return []string{}, err
		}
	}

	// Get tree.
	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	// Get file.
	entry, err := tree.EntryByPath(filePath)
	if err != nil {
		return nil, err
	}

	blob, err := repo.LookupBlob(entry.Id)
	if err != nil {
		return nil, err
	}

	// Get content.
	content := blob.Contents()

	// Split lines.
	res := strings.Split(string(content[:]), "\n")

	return res, nil
}
