package database

import (
	"fmt"

	"github.com/guregu/dynamo"
	"github.com/josephnp732/Stocks-GraphQL/graph/model"
)

// GetStocksByRange returns stock data between date range
func GetStocksByRange(input model.StockInput) ([]*model.Stock, error) {

	var result []*model.Stock

	if input.FromDate > input.ToDate {
		return nil, fmt.Errorf("Dates of out range. Please check dates")
	}

	query := Table.Get("stock_symbol", input.StockSymbol).Range("date", dynamo.Between, input.FromDate, input.ToDate).All(&result)
	if query != nil {
		return nil, fmt.Errorf(query.Error())
	}

	return result, nil
}

// GetStocksByTicker returns the stocks by a ticker_symbol
func GetStocksByTicker(symbol string) ([]*model.Stock, error) {
	var result []*model.Stock

	query := Table.Get("stock_symbol", symbol).All(&result)
	if query != nil {
		return nil, fmt.Errorf(query.Error())
	}

	return result, nil
}
