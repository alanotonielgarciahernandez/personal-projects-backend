// graphQLRequest.go
// Request model used to validate user.

package models

// GraphQLRequest, HTTP Request Body model.
type GraphQLRequest struct {
	Query string `json:"query,omitempty"`
}
