// getTree.go
// Project tree handler.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/git"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	git2go "github.com/libgit2/git2go/v34"
)

// Gets project tree.
func GetTree() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validate project header.
		project := ctx.Request.Header.Get("Project")
		if len(project) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing project parameter.",
			})
			return
		}

		// Validate path header.
		path := ctx.Request.Header.Get("Path")
		/*
			if len(path) < 1 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Missing path parameter.",
				})
				return
			}
		*/

		// Reference header is not required.
		reference := ctx.Request.Header.Get("Reference")

		// Open repository.
		repo, _ := git2go.OpenRepository(os.Getenv("REPOSITORIES_DIR") + project + ".git")

		// Search for reference hash.
		branch, err := repo.LookupBranch(reference, git2go.BranchAll)
		if err == nil {
			reference = branch.Target().String()
		} else {
			tag, err := repo.References.Lookup("refs/tags/" + reference)
			if err == nil {
				reference = tag.Target().String()
			}
		}

		// Get tree.
		res, err := git.GetTree(repo, reference, path)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Request response.
		ctx.JSON(http.StatusAccepted, res)
	}
}
