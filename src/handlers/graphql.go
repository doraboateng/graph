// TODO: move this into a graph service.
package handlers

import (
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

// Define a schema string:
const schemaString = `
	# Define what the schema is capable of:
	schema {
		query: Query
	}
	# Define what the queries are capable of:
	type Query {
		# Generic greeting, e.g. "Hello, world!":
		greet: String!
	}
`

// RootResolver ...
type RootResolver struct{}

// Greet ...
func (*RootResolver) Greet() string {
	return "Hello, world!"
}

// Schema ...
var Schema = graphql.MustParseSchema(schemaString, &RootResolver{})

// ClientQuery ...
type ClientQuery struct {
	Operation string                 // Operation name.
	Query     string                 // Query string.
	Variables map[string]interface{} // Query variables (untyped).
}

// GraphHandler ...
func GraphHandler(writer http.ResponseWriter, request *http.Request) {
	var params struct {
		Query     string                 `json:"query"`
		Operation string                 `json:"operationName"`
		Variables map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(request.Body).Decode(&params); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := Schema.Exec(request.Context(), params.Query, params.Operation, params.Variables)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Write(jsonResponse)
}
