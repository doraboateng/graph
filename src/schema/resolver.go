package schema

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/kwcay/boateng-graph-service/src/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver ...
type Resolver struct {
	alphabets   []*generated.Alphabet
	expressions []*generated.Expression
	languages   []*generated.Language
}

func (resolver *queryResolver) Alphabets(ctx context.Context) ([]*generated.Alphabet, error) {
	return resolver.alphabets, nil
}

func (resolver *queryResolver) Expressions(ctx context.Context) ([]*generated.Expression, error) {
	return resolver.expressions, nil
}

func (resolver *queryResolver) Languages(ctx context.Context) ([]*generated.Language, error) {
	return resolver.languages, nil
}
