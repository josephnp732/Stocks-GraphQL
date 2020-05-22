package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	stocks "github.com/josephnp732/Stocks-GraphQL/database"

	"github.com/josephnp732/Stocks-GraphQL/graph/generated"
	"github.com/josephnp732/Stocks-GraphQL/graph/model"
)

// AddStock creates a new Stock in the Database
func (r *mutationResolver) AddStock(ctx context.Context, input model.NewStock) (*model.Stock, error) {
	newStock, err := stocks.AddStock(input)
	r.stocks = append(r.stocks, *newStock)
	return newStock, err
}

// Get a stock from the Database
func (r *queryResolver) GetStocks(ctx context.Context, stockSymbol string) ([]*model.Stock, error) {
	return stocks.GetStocksByTicker(stockSymbol)
}

func (r *queryResolver) GetStocksByRange(ctx context.Context, input model.StockInput) ([]*model.Stock, error) {
	return stocks.GetStocksByRange(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
