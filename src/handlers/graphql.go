package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kwcay/boateng-graph-service/src/generated"
	"github.com/kwcay/boateng-graph-service/src/schema"
)

// GraphHandler ...
func GraphHandler(writer http.ResponseWriter, request *http.Request) {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &schema.Resolver{}}),
	)

	srv.ServeHTTP(writer, request)
}
