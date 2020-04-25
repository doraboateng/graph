package resolvers

import (
	"context"

	"github.com/kwcay/boateng-graph-service/src/generated"
)

type expressionResolver struct{ *Resolver }

func (r *queryResolver) Expression(ctx context.Context, code string) (*generated.Expression, error) {
	panic("not implemented")
}

func (r *queryResolver) Expressions(ctx context.Context) ([]*generated.Expression, error) {
	panic("not implemented")
}
