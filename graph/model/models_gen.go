package model

type NewStock struct {
	Exchange           string `json:"exchange"`
	StockSymbol        string `json:"stock_symbol"`
	Date               string `json:"date"`
	StockPriceOpen     string `json:"stock_price_open"`
	StockPriceHigh     string `json:"stock_price_high"`
	StockPriceLow      string `json:"stock_price_low"`
	StockPriceClose    string `json:"stock_price_close"`
	StockVolume        string `json:"stock_volume"`
	StockPriceAdjClose string `json:"stock_price_adj_close"`
}

type Stock struct {
	Exchange           string `json:"exchange" dynamo:"exchange"`
	StockSymbol        string `json:"stock_symbol" dynamo:"stock_symbol"`
	Date               string `json:"date" dynamo:"date"`
	StockPriceOpen     string `json:"stock_price_open" dynamo:"stock_price_open"`
	StockPriceHigh     string `json:"stock_price_high" dynamo:"stock_price_high"`
	StockPriceLow      string `json:"stock_price_low" dynamo:"stock_price_low"`
	StockPriceClose    string `json:"stock_price_close" dynamo:"stock_price_close"`
	StockVolume        string `json:"stock_volume" dynamo:"stock_volume"`
	StockPriceAdjClose string `json:"stock_price_adj_close" dynamo:"stock_price_adj_close"`
}

type StockInput struct {
	StockSymbol string `json:"stock_symbol"`
	FromDate    string `json:"from_date"`
	ToDate      string `json:"to_date"`
}
