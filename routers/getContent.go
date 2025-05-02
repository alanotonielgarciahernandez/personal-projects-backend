// getContent.go
// File content handler.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/git"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handles the selected file content.
func GetContent() gin.HandlerFunc {
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
		if len(path) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing path parameter.",
			})
			return
		}

		// Reference header is not required.
		reference := ctx.Request.Header.Get("Reference")

		// Get file content.
		res, err := git.GetContent(project, path, reference)
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
