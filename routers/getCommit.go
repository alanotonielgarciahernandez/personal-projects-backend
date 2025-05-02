// getCommit.go
// Commit data handler.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/git"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handles the selected commit data.
func GetCommit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validate project header.
		project := ctx.Request.Header.Get("Project")
		if len(project) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing project parameter.",
			})
			return
		}

		// Validate reference header.
		reference := ctx.Request.Header.Get("Reference")
		if len(reference) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing reference parameter.",
			})
			return
		}

		// Get commit data.
		res, err := git.GetCommit(project, reference)
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
