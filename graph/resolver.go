package graph

import "github.com/josephnp732/Stocks-GraphQL/graph/model"

// Resolver is the resolver for the system
type Resolver struct {
	stocks []*model.Stock
}
