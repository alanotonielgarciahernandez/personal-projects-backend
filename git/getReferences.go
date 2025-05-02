// getReferences.go
// Gets a project references data.

package git

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"errors"
	"strings"

	git "github.com/libgit2/git2go/v34"
)

// Gets a project references data.
// Returns reference list, branch count, commit count, tag count and head commit.
func GetReferences(project *git.Repository, projectName string) ([]models.ReferenceInput, int, int, int, models.ReferenceInput, error) {
	// Variables.
	var branchCount int = 0
	var commitCount int = 0
	var tagCount int = 0
	var res = []models.ReferenceInput{}
	var headCommit models.ReferenceInput

	head, err := project.Head()
	if err != nil {
		return res, branchCount, commitCount, tagCount, headCommit, errors.New("while getting repository head, " + err.Error())
	}

	iter, err := project.NewReferenceIterator()
	if err != nil {
		return res, branchCount, commitCount, tagCount, headCommit, errors.New("while getting reference iterator, " + err.Error())
	}
	defer iter.Free()

	walker, err := project.Walk()
	if err != nil {
		return res, branchCount, commitCount, tagCount, headCommit, errors.New("while getting repository walker, " + err.Error())
	}
	defer walker.Free()

	// Walk through references.
	for {
		ref, err := iter.Next()
		if err != nil {
			break
		}

		commit, err := project.LookupCommit(ref.Target())
		if err != nil {
			continue
		}

		// If current reference is a branch, add branch data into reference list.
		if ref.IsBranch() {
			// Increase branch count.
			branchCount++
			isHead, _ := ref.Branch().IsHead()
			walker.Push(ref.Target())

			// Get ahead/behind data.
			ahead, behind, err := ref.Owner().AheadBehind(ref.Target(), head.Target())
			if err != nil {
				return res, branchCount, commitCount, tagCount, headCommit, errors.New("while getting ahead behind data, " + err.Error())
			}

			// Push into reference list.
			res = append(res, models.ReferenceInput{
				Hash:    ref.Target().String(),
				Project: projectName,
				Name:    strings.Split(ref.Name(), "/")[2],
				Type:    "branch",
				Date:    commit.Author().When,
				Default: &isHead,
				Aheadbehind: &model.AheadBehind{
					Ahead:  ahead,
					Behind: behind,
				},
			})

			// If current reference is the default branch, add it to the respective model.
			if isHead {
				headCommit = models.ReferenceInput{
					Hash:    ref.Target().String(),
					Project: projectName,
					Name:    strings.Split(ref.Name(), "/")[2],
					Type:    "branch",
					Date:    commit.Author().When,
					Default: &isHead,
					Aheadbehind: &model.AheadBehind{
						Ahead:  ahead,
						Behind: behind,
					},
				}
			}
		} else if ref.IsTag() {
			// Else if current reference is a tag, add tag data into reference list.

			// Increase tag count.
			tagCount++

			res = append(res, models.ReferenceInput{
				Hash:    ref.Target().String(),
				Project: projectName,
				Name:    strings.Split(ref.Name(), "/")[2],
				Type:    "tag",
				Date:    commit.Author().When,
				Message: strings.Split(strings.Trim(commit.Message(), "\n"), "\n\n"),
			})
		}
	}

	// Sort commits by time.
	walker.Sorting(git.SortTime)

	// Walk through commits.
	walker.Iterate(func(commit *git.Commit) bool {
		// Increase commit count.
		commitCount++

		// Add commit data into reference list.
		res = append(res, models.ReferenceInput{
			Hash:    commit.Id().String(),
			Project: projectName,
			Type:    "commit",
			Date:    commit.Author().When,
			Message: strings.Split(strings.Trim(commit.Message(), "\n"), "\n\n"),
		})
		return true
	})

	return res, branchCount, commitCount, tagCount, headCommit, nil
}
