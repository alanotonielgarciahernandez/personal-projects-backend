// getArchive.go
// Archive function handler.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/git"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handles archive function.
func GetArchive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validate Project header.
		project := ctx.Request.Header.Get("Project")
		if len(project) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing project parameter.",
			})
			return
		}

		// Validate Reference header.
		reference := ctx.Request.Header.Get("Reference")
		if len(project) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing reference parameter.",
			})
			return
		}

		// Validate Extension header.
		extension := ctx.Request.Header.Get("Extension")
		if len(project) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing extension parameter.",
			})
			return
		}

		// Get project compressed file.
		res, err := git.GetArchive(project, reference, extension)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Request response.
		ctx.Header("Content-Type", "application/zip")
		ctx.Header("Content-Disposition", "attachment; filename=repo-"+reference+".zip")
		ctx.File(res)
	}
}
