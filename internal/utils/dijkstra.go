package utils

import (
	"binance-trade-bot/internal/binance"
	"github.com/sirupsen/logrus"
)

func CalculateBestPrice(symbolsToSymbols map[string]map[string]*binance.ExchangeSymbol, currentSymbol string) (float64, []string) {
	currenciesForCheck := make(map[string]bool)
	currenciesForCheck[currentSymbol] = true

	pathes := make(map[string]float64)
	p := make(map[string]string)
	pathes[currentSymbol] = 1.0
	for currencySymbol := range currenciesForCheck {
		currenciesForCheck[currentSymbol] = false
		delete(currenciesForCheck, currentSymbol)

		for _, currency := range symbolsToSymbols[currencySymbol] {
			if _, ok := pathes[currency.QuoteAsset]; !ok || pathes[currency.QuoteAsset] < pathes[currency.BaseAsset]*currency.PriceApproximateInQuote {
				assetCheck := currency.BaseAsset
				for {
					if _, ok := p[assetCheck]; !ok {
						break
					}

					assetCheck = p[assetCheck]
					if assetCheck == currency.QuoteAsset {
						break
					}
				}
				if assetCheck == currency.QuoteAsset {
					continue
				}

				pathes[currency.QuoteAsset] = pathes[currency.BaseAsset] * currency.PriceApproximateInQuote
				p[currency.QuoteAsset] = currency.BaseAsset
			}

			currenciesForCheck[currency.QuoteAsset] = true
		}
	}

	var pathesToCurrency []string
	maxPrice := 0.0
	for _, items := range symbolsToSymbols {
		for _, currency := range items {
			if currency.QuoteAsset != currentSymbol {
				continue
			}

			steps := 0
			cur0 := currency.BaseAsset
			for {
				if _, ok := p[cur0]; !ok {
					break
				}

				cur0 = p[cur0]
				steps++
			}
			if steps == 1 {
				continue
			}

			price := currency.PriceApproximateInQuote * pathes[currency.BaseAsset]
			if price <= 1.0 {
				continue
			}
			if price >= maxPrice {
				maxPrice = price
				pathesToCurrency = []string{currency.QuoteAsset, currency.BaseAsset}
			}

			cur := currency.BaseAsset
			logrus.Infof("from %s to %s is %f", currency.BaseAsset, currency.QuoteAsset, symbolsToSymbols[currency.BaseAsset][currency.QuoteAsset].PriceApproximateInQuote)

			for {
				if _, ok := p[cur]; !ok {
					break
				}

				logrus.Infof("from %s to %s is %f", p[cur], cur, symbolsToSymbols[p[cur]][cur].PriceApproximateInQuote)
				cur = p[cur]
				if maxPrice == price {
					pathesToCurrency = append(pathesToCurrency, cur)
				}
			}
			logrus.Infof("price to %s is %f %v", currency.QuoteAsset, price, pathesToCurrency)
		}
	}

	logrus.Infof("maxPrice is %f, path", maxPrice)
	return maxPrice, pathesToCurrency
}

func Dijkstra1(symbolsToSymbols map[string]map[string]*binance.ExchangeSymbol, currentSymbol string) float64 {
	currencyToVertex := make(map[string]int)
	// обязательно оставить двумя циклами, тк так кодируются порядок заполнения
	for baseAsset := range symbolsToSymbols {
		if _, ok := currencyToVertex[baseAsset]; !ok {
			currencyToVertex[baseAsset] = len(currencyToVertex)
		}
	}

	for _, symbols := range symbolsToSymbols {
		for quoteAsset := range symbols {
			if _, ok := currencyToVertex[quoteAsset]; !ok {
				currencyToVertex[quoteAsset] = len(currencyToVertex)
			}
		}
	}

	type VertexData struct {
		toVertex int
		weight   float64
	}
	vertex := make([][]VertexData, len(symbolsToSymbols))
	for _, symbols := range symbolsToSymbols {
		for _, currency := range symbols {
			vertex[currencyToVertex[currency.BaseAsset]] = append(vertex[currencyToVertex[currency.BaseAsset]], VertexData{
				toVertex: currencyToVertex[currency.QuoteAsset],
				weight:   currency.PriceApproximateInQuote,
			})
		}
	}

	current := currencyToVertex[currentSymbol]

	// algo
	type Edge struct {
		from   int
		to     int
		weight float64
	}
	var edges []Edge
	for i := range vertex {
		for j := range vertex[i] {
			edges = append(edges, Edge{
				from:   i,
				to:     vertex[i][j].toVertex,
				weight: vertex[i][j].weight,
			})
		}
	}

	d := make(map[int]float64)
	p := make(map[int]int)
	d[current] = 1

	for {
		var any bool
		for _, edge := range edges {
			from := edge.from
			to := edge.to
			if _, ok := d[from]; !ok {
				continue
			}

			finalWeight := d[from] + edge.weight
			if _, ok := d[to]; !ok || finalWeight < d[to] {
				d[to] = finalWeight
				p[to] = from
				any = true
			}
		}

		if !any {
			break
		}
	}

	num := current
	num = currencyToVertex["TRX"]
	/*for key, value := range p {
		if value == num {
			var from string
			var to string
			for name, index := range currencyToVertex {
				if index == num {
					from = name
				}
				if index == p[num] {
					to = name
				}
			}

			logrus.Infof("from %s[%d] to %s[%d]", from, num, to, p[num])
			num = p[num]
		}
	}*/
	num = currencyToVertex["USDT"]
	// num = currencyToVertex["BTC"]
	num = currencyToVertex["CTSI"]
	logrus.Infof("path to %d is %f", num, d[num])
	for {
		if _, ok := p[num]; !ok {
			break
		}
		var from string
		var to string
		for name, index := range currencyToVertex {
			if index == num {
				from = name
			}
			if index == p[num] {
				to = name
			}
		}

		logrus.Infof("from %s[%d] to %s[%d]", from, num, to, p[num])
		num = p[num]
	}

	return d[current]
}

func Dijkstra2(symbolsToSymbols map[string]map[string]*binance.ExchangeSymbol, currentSymbol string) float64 {
	INF := -100000.0
	EMPTY_INDEX := ""

	d := make(map[string]float64)
	p := make(map[string]string)
	u := make(map[string]bool)
	for symbolKey := range symbolsToSymbols {
		d[symbolKey] = INF
	}
	d[currentSymbol] = 1

	for _ = range symbolsToSymbols {
		v := EMPTY_INDEX
		for baseAsset := range symbolsToSymbols {
			if !u[baseAsset] && (v == "" || d[baseAsset] > d[v]) {
				v = baseAsset
			}
		}
		if d[v] == INF {
			break
		}
		u[v] = true

		for quoteAsset, symbol := range symbolsToSymbols[v] {
			priceNewPath := d[v] * symbol.PriceApproximateInQuote
			if priceNewPath > d[quoteAsset] {
				d[quoteAsset] = priceNewPath
				p[quoteAsset] = v
			}
		}
	}

	q1 := p[currentSymbol]
	q2 := p[q1]
	q3 := p[q2]
	q4 := p[q3]
	_ = q4

	return d[currentSymbol]
}
