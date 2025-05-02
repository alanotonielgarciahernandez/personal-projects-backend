// getImage.go
// Image handler.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/cfg"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Handles the selected file content.
func GetImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filename := ctx.Param("filename")

		imagePath := filepath.Join(cfg.UPLOADS_FOLDER, filename)

		ctx.File(imagePath)
	}
}
