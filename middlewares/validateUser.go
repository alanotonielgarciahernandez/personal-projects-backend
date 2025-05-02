// validateUser.go
// Middleware valitating administrator user when making a mutation request.

package middlewares

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Validate mutation request.
func isMutation(query string) bool {
	trimmedQuery := strings.TrimSpace(query)
	return strings.HasPrefix(trimmedQuery, "mutation")
}

// Validates the authorization header for mutation operations.
func ValidateUser() gin.HandlerFunc {

	AUTHORIZATION_KEY := os.Getenv("AUTHORIZATION_KEY")

	if AUTHORIZATION_KEY == "" {
		log.Fatal("Missing environment variable AUTHORIZATION_KEY")
	}

	return func(ctx *gin.Context) {
		var req models.GraphQLRequest

		contentType := ctx.Request.Header.Get("Content-Type")
		buf, _ := io.ReadAll(ctx.Request.Body)

		if strings.HasPrefix(contentType, "application/json") {

			err := json.Unmarshal(buf, &req)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Invalid request body",
				})
				return
			}
		} else {
			// Reconstruct the multipart form to read the "operations" field
			boundary := strings.Split(contentType, "boundary=")
			if len(boundary) < 2 {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing multipart boundary"})
				return
			}

			mr := multipart.NewReader(bytes.NewReader(buf), boundary[1])
			for {
				part, err := mr.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					break
				}

				if part.FormName() == "operations" {
					operationsData, _ := io.ReadAll(part)
					err := json.Unmarshal(operationsData, &req)
					if err != nil {
						ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"error": "Invalid request body",
						})
						return
					}
					break
				}
			}
		}

		if isMutation(req.Query) {
			if ctx.Request.Header.Get("Authorization") != AUTHORIZATION_KEY {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user",
				})
				return
			}
		}

		// Restore the request body for gqlgen
		ctx.Request.Body = io.NopCloser(bytes.NewReader(buf))

		ctx.Next()
	}
}
