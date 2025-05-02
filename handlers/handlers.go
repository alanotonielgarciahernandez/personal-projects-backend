// handlers.go
// Defines endpoints.

package handlers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/middlewares"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Sets the engine configuration and endpoints.
func Managers() {
	router := gin.Default()

	// CORS configuration.
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Project", "Reference", "File", "Path"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.SetTrustedProxies(nil)

	router.Use(cors.New(config))

	// GraphQL endpoint.
	router.Any("/query", middlewares.CheckDB(), middlewares.ValidateUser(), routers.GraphQL())

	// Custom endpoints.
	router.GET("/get-archive", middlewares.CheckDB(), routers.GetArchive())
	router.GET("/get-commit", middlewares.CheckDB(), routers.GetCommit())
	router.GET("/get-content", middlewares.CheckDB(), routers.GetContent())
	router.GET("/get-files", middlewares.CheckDB(), routers.GetTree())
	router.GET("/get-importantfiles", middlewares.CheckDB(), routers.GetImportantFiles())
	router.GET("/uploads/:filename", middlewares.CheckDB(), routers.GetImage())

	// Run HTTP Server on port 3000.
	router.Run(":3000")
}
