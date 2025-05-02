// checkDB.go
// Middleware checking the connection to the database.

package middlewares

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Checks if the database is active.
func CheckDB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if db.CheckDBConnection() != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Lost connection with database.",
			})
			return
		}

		ctx.Next()
	}
}
