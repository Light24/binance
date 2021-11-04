package binance

type ExchangeSymbol struct {
	PriceApproximateInQuote float64
	BaseAsset               string
	QuoteAsset              string
	Filters                 []map[string]interface{}
}
