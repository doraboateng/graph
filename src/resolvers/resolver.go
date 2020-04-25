package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/kwcay/boateng-graph-service/src/generated"
)

// Resolver ...
type Resolver struct {
	Dgraph *dgo.Dgraph
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
