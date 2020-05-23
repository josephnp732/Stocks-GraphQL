package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gobuffalo/x/randx"

	stocks "github.com/josephnp732/Stocks-GraphQL/database"
	"github.com/josephnp732/Stocks-GraphQL/graph/generated"
	"github.com/josephnp732/Stocks-GraphQL/graph/model"
)

var stockAddedChannel map[string]chan *model.Stock

func init() {
	stockAddedChannel = map[string]chan *model.Stock{}
}

func (r *mutationResolver) AddStock(ctx context.Context, input model.NewStock) (*model.Stock, error) {
	newStock, err := stocks.AddStock(input)
	r.stocks = append(r.stocks, newStock)

	for _, observer := range stockAddedChannel {
		observer <- newStock
	}

	return newStock, err
}

func (r *queryResolver) GetStocks(ctx context.Context, stockSymbol string) ([]*model.Stock, error) {
	return stocks.GetStocksByTicker(stockSymbol)
}

func (r *queryResolver) GetStocksByRange(ctx context.Context, input model.StockInput) ([]*model.Stock, error) {
	return stocks.GetStocksByRange(input)
}

func (r *subscriptionResolver) StockAdded(ctx context.Context) (<-chan *model.Stock, error) {
	id := randx.String(8)

	stockEvent := make(chan *model.Stock, 1)
	go func() {
		<-ctx.Done()
	}()
	stockAddedChannel[id] = stockEvent

	return stockEvent, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
