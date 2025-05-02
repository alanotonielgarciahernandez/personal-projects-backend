package graph

import "alan-otoniel-garcia-hernandez-personal-projects/backend/graph/model"

//go:generate go get github.com/99designs/gqlgen/graphql/handler/extension@v0.17.66
//go:generate go get github.com/99designs/gqlgen/graphql/handler/lru@v0.17.66
//go:generate go get github.com/99designs/gqlgen/graphql/handler/transport@v0.17.66
//go:generate go get github.com/99designs/gqlgen@v0.17.66
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	projects []model.Project
	issues   []model.ProjectIssue
}
