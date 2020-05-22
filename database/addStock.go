package database

import (
	"fmt"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/josephnp732/Stocks-GraphQL/graph/model"
)

// Sample date format
const layoutISO string = "2006-01-02"

// AddStock adds a new stock to dynamoDB
func AddStock(input model.NewStock) (model.Stock, error) {

	t := time.Now().UTC()
	date := t.Format(layoutISO)

	newStock := model.Stock{
		Exchange:           "NYSE",
		StockSymbol:        input.StockSymbol,
		Date:               date,
		StockPriceAdjClose: input.StockPriceAdjClose,
		StockPriceClose:    input.StockPriceClose,
		StockPriceHigh:     input.StockPriceHigh,
		StockPriceLow:      input.StockPriceLow,
		StockPriceOpen:     input.StockPriceOpen,
	}

	query := Table.Put(&newStock).Run()
	if query != nil {
		fmt.Print(query.Error())
		return model.Stock{}, gqlerror.Errorf("Unable to add Stock to the Database")
	}

	return newStock, nil
}
