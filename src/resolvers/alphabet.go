package resolvers

import (
	"context"

	"github.com/kwcay/boateng-graph-service/src/generated"
)

type alphabetResolver struct{ *Resolver }

func (r *queryResolver) Alphabet(ctx context.Context, code string) (*generated.Alphabet, error) {
	panic("not implemented")
}

func (r *queryResolver) Alphabets(ctx context.Context) ([]*generated.Alphabet, error) {
	panic("not implemented")
}
