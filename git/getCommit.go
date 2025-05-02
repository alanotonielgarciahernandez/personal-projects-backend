// getCommit.go
// Gets commit data.

package git

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"errors"
	"os"
	"strings"
	"time"

	git "github.com/libgit2/git2go/v34"
)

// Gets commit data.
func GetCommit(project string, reference string) (*models.Commit, error) {
	// Response model.
	var res *models.Commit = &models.Commit{
		Diff:         make(map[string][]string),
		LinesChanged: models.LinesChanged{Added: 0, Deleted: 0},
	}

	// Variables.
	var parentCommit *git.Commit
	var parentId string = ""
	var parentTree *git.Tree

	REPOSITORIES_DIR := os.Getenv("REPOSITORIES_DIR")

	if REPOSITORIES_DIR == "" {
		return nil, errors.New("missing environment variable REPOSITORIES_DIR")
	}

	repo, err := git.OpenRepository(REPOSITORIES_DIR + project + ".git")
	if err != nil {
		return nil, err
	}

	ref, err := git.NewOid(reference)
	if err != nil {
		return nil, err
	}

	// Get commit object.
	commit, err := repo.LookupCommit(ref)
	if err != nil {
		return nil, err
	}
	defer commit.Free()

	commitTree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	defer commitTree.Free()

	// Get commit parent.
	if commit.ParentCount() > 0 {
		parentCommit = commit.Parent(0)
		defer parentCommit.Free()

		parentId = parentCommit.Id().String()

		parentTree, err = parentCommit.Tree()
		if err != nil {
			return nil, err
		}
		defer parentTree.Free()
	}

	// Get diff object.
	diffOptions, _ := git.DefaultDiffOptions()
	diff, err := repo.DiffTreeToTree(parentTree, commitTree, &diffOptions)
	if err != nil {
		return nil, err
	}
	defer diff.Free()

	// Format diff data.
	err = diff.ForEach(func(delta git.DiffDelta, progress float64) (git.DiffForEachHunkCallback, error) {
		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			res.Diff[delta.NewFile.Path] = append(res.Diff[delta.NewFile.Path], hunk.Header)
			return func(line git.DiffLine) error {
				if line.Origin.String() == "Addition" {
					res.LinesChanged.Added++
				} else if line.Origin.String() == "Deletion" {
					res.LinesChanged.Deleted++
				}
				res.Diff[delta.NewFile.Path] = append(res.Diff[delta.NewFile.Path], line.Origin.String()+" "+strings.Split(line.Content, "\n")[0])
				return nil
			}, nil
		}, nil
	}, git.DiffDetailLines)
	if err != nil {
		return nil, err
	}

	res.Branch = ""
	res.Date = commit.Author().When.Format(time.RFC1123)
	res.Message = strings.Split(strings.Trim(commit.Message(), "\n"), "\n\n")
	res.Parent = parentId

	return res, nil
}
