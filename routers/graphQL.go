// graphQL.go
// Endpoint for GraphQL Server.

package routers

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
)

// GraphQL Handler.
func GraphQL() gin.HandlerFunc {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	return func(ctx *gin.Context) {

		// Run GraphQL Handler if request body is valid.
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
