package strategies

import (
	"binance-trade-bot/internal"
	"binance-trade-bot/internal/binance"
	"binance-trade-bot/internal/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type exchangeChain struct {
	binanceManager binance.Manager
	config         internal.AppConfig

	balances map[string]float64
}

func NewExchangeChain(binanceManager binance.Manager, config internal.AppConfig) Strategy {
	return &exchangeChain{
		binanceManager,
		config,
		nil,
	}
}

// Scout - Scout for potential jumps from the current coin to another coin
func (strategy *exchangeChain) Scout() error {
	account, err := strategy.binanceManager.GetAccount()
	logrus.Info(account)

	strategy.balances = make(map[string]float64)
	for _, item := range account.Balances {
		balanceFree, err := strconv.ParseFloat(item.Free, 64)
		if err != nil {
			return err
		}

		if balanceFree == 0 {
			continue
		}

		strategy.balances[item.Asset] = balanceFree
	}
	logrus.Info(strategy.balances)

	currentSymbol, err := strategy.binanceManager.GetCurrentCoin()
	if err != nil {
		return err
	}

	currentCoinPrice, err := strategy.binanceManager.GetTickerPrice(currentSymbol + strategy.config.Bridge)
	if err != nil {
		return err
	}

	return strategy.jumpToBestCoin(currentSymbol, currentCoinPrice)
}

func (strategy *exchangeChain) jumpToBestCoin(currentSymbol string, currentCoinPrice float64) error {
	ratios, err := strategy.getRatios(currentSymbol, currentCoinPrice)
	if err != nil {
		return err
	}

	var bestRatio *ratioPair
	for _, ratio := range ratios {
		if bestRatio == nil || bestRatio.ratio < ratio.ratio {
			bestRatio = &ratio
		}
	}
	if bestRatio == nil {
		err := errors.Errorf("jumpToBestCoin: best ration not found")
		return err
	}

	return strategy.transactionThroughBridge(bestRatio)
}

// getRatios - Given a coin, get the current price ratio for every other enabled coin
func (strategy *exchangeChain) getRatios(currentSymbol string, currentCoinPrice float64) ([]ratioPair, error) {
	supportedSymbols, err := strategy.binanceManager.GetSupportedSymbols()
	if err != nil {
		return nil, err
	}

	optionalCoinPrices, err := strategy.binanceManager.GetTickerPrices()
	if err != nil {
		return nil, err
	}

	for baseSymbol, supportedSymbol := range supportedSymbols {
		for quoteSymbol, exchange := range supportedSymbol {
			priceApproximate := optionalCoinPrices[baseSymbol+quoteSymbol]
			if priceApproximate != 0 {
				priceApproximate = priceApproximate - priceApproximate*0.075 // commission

				exchange.PriceApproximateInQuote = priceApproximate

				if supportedSymbols[quoteSymbol] == nil {
					supportedSymbols[quoteSymbol] = make(map[string]*binance.ExchangeSymbol)
				}
				if supportedSymbols[quoteSymbol][baseSymbol] == nil {
					supportedSymbols[quoteSymbol][baseSymbol] = &binance.ExchangeSymbol{
						BaseAsset:               quoteSymbol,
						QuoteAsset:              baseSymbol,
						PriceApproximateInQuote: 1 / priceApproximate,
					}
				}
			}
		}
	}

	/*var maxRation float64
	for symbol := range supportedSymbols {
		ratio := utils.CalculateBestPrice(supportedSymbols, symbol)
		maxRation = math.Max(maxRation, ratio)
	}
	logrus.Infof("maxRation is %f", maxRation)
	*/

	ratio, pathesToCurrency := utils.CalculateBestPrice(supportedSymbols, currentSymbol)
	logrus.Infof("path to %s is %f %v", currentSymbol, ratio, pathesToCurrency)
	if ratio <= 1.0 {
		return nil, errors.New("Not found ratio")
	}

	price := 1.0
	for i := range pathesToCurrency {
		if pathesToCurrency[0] == pathesToCurrency[i] && i != 0 {
			logrus.Infof("price %f vs real price %f", ratio, price)
			if price < 1.000 {
				return nil, errors.New("Not found ratio")
			}
			break
		}

		fromCoinPrice, err := strategy.binanceManager.GetTickerPrice(pathesToCurrency[i] + pathesToCurrency[i+1])
		if err != nil {
			fromCoinPrice, err = strategy.binanceManager.GetTickerPrice(pathesToCurrency[i+1] + pathesToCurrency[i])
			logrus.Infof("PRICE from %s to %s is %f", pathesToCurrency[i+1], pathesToCurrency[i], fromCoinPrice)
			price *= 1 / fromCoinPrice
		} else {
			logrus.Infof("PRICE from %s to %s is %f", pathesToCurrency[i], pathesToCurrency[i+1], fromCoinPrice)
			price *= fromCoinPrice
		}
		// fromCoinPrice, err := strategy.binanceManager.GetOpenOrder(pathesToCurrency[i] + pathesToCurrency[i + 1])
		if err != nil {
			return nil, err
		}

	}

	return nil, errors.New("Not found ratio")
}

// Jump from the source coin to the destination coin through bridge coin
func (strategy *exchangeChain) transactionThroughBridge(bestRatio *ratioPair) error {
	balance, err := strategy.binanceManager.GetCurrencyBalance(bestRatio.origin)
	if err != nil {
		return err
	}

	fromCoinPrice, err := strategy.binanceManager.GetTickerPrice(bestRatio.origin + strategy.config.Bridge)
	if err != nil {
		return err
	}

	minNotional, err := strategy.binanceManager.GetMinNotional(bestRatio.origin, bestRatio.target)
	if err != nil {
		return err
	}

	canSell := balance*fromCoinPrice > minNotional
	if canSell {
		resp, err := strategy.binanceManager.SellAlt(bestRatio.origin, strategy.config.Bridge)
		if err != nil {
			return err
		} else if resp == nil {
			err := errors.Errorf("Couldn't sell, going back to scouting mode..")
			logrus.Error(err)
			return err
		}
	}

	resp, err := strategy.binanceManager.BuyAlt(bestRatio.target, strategy.config.Bridge)
	if err != nil {
		return err
	} else if resp == nil {
		err := errors.Errorf("Couldn't buy, going back to scouting mode...")
		logrus.Error(err)
		return err
	}

	return nil
}
