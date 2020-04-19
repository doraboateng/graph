package schema

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/kwcay/boateng-graph-service/src/generated"
)

// Query returns generated.QueryResolver implementation.
func (resolver *Resolver) Query() generated.QueryResolver {
	return &queryResolver{resolver}
}

type queryResolver struct{ *Resolver }
