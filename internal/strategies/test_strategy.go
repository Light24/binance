package strategies

import (
	"binance-trade-bot/internal"
	"binance-trade-bot/internal/binance"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"strconv"
)

type testStrategy struct {
	binanceManager binance.Manager
	config         internal.AppConfig

	balances map[string]float64
}

func NewTestStrategy(binanceManager binance.Manager, config internal.AppConfig) Strategy {
	return &testStrategy{
		binanceManager,
		config,
		nil,
	}
}

// Scout - Scout for potential jumps from the current coin to another coin
func (strategy *testStrategy) Scout() error {
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

func (strategy *testStrategy) jumpToBestCoin(currentSymbol string, currentCoinPrice float64) error {
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

	return strategy.jumpBestCoin(bestRatio)
}

// getRatios - Given a coin, get the current price ratio for every other enabled coin
func (strategy *testStrategy) getRatios(currentSymbol string, currentCoinPrice float64) ([]ratioPair, error) {
	supportedSymbols, err := strategy.binanceManager.GetSupportedSymbols()
	if err != nil {
		return nil, err
	}

	optionalCoinPrices, err := strategy.binanceManager.GetTickerPrices()
	if err != nil {
		return nil, err
	}

	tradeFees, err := strategy.binanceManager.GetTradeFees()
	if err != nil {
		return nil, err
	}

	usingBnbForFees, err := strategy.binanceManager.GetUsingBnbForFees()
	if err != nil {
		return nil, err
		// usingBnbForFees = false
	}

	var ratios []ratioPair
	for baseAsset, supportedSymbol := range supportedSymbols {
		if baseAsset != currentSymbol {
			continue
		}

		var exchangeSymbol *binance.ExchangeSymbol
		for _, item := range supportedSymbol {
			if item.QuoteAsset == strategy.config.Bridge {
				exchangeSymbol = item
				continue
			}
		}
		if exchangeSymbol == nil {
			continue
		}

		optionalCoinPrice := optionalCoinPrices[baseAsset+exchangeSymbol.QuoteAsset]
		if optionalCoinPrice == 0 {
			continue
		}

		baseFee := tradeFees[currentSymbol+exchangeSymbol.QuoteAsset]
		if baseFee.Equal(decimal.Decimal{}) {
			continue
		}

		currentBalance := strategy.balances[currentSymbol]
		targetBalance := strategy.balances[exchangeSymbol.QuoteAsset]
		transactionFee, err := strategy.binanceManager.GetFee(currentSymbol, exchangeSymbol.QuoteAsset, currentBalance, targetBalance, baseFee, usingBnbForFees, optionalCoinPrice, false)
		if err != nil {
			continue
		}

		// https://github.com/edeng23/binance-trade-bot/blob/d76e4da7b5b41a925cf06d3c33c34b26db191932/binance_trade_bot/auto_trader.py#L131
		// https://github.com/edeng23/binance-trade-bot/blob/d76e4da7b5b41a925cf06d3c33c34b26db191932/binance_trade_bot/auto_trader.py#L71
		coinOptCoinRatio := currentCoinPrice / optionalCoinPrice
		ratio := coinOptCoinRatio*(1-transactionFee*strategy.config.ScoutMultiplier) - (currentCoinPrice / optionalCoinPrice)
		ratios = append(ratios, ratioPair{currentSymbol, baseAsset, ratio})

		logrus.Infof("Transfer from %s to %s ratio is %f", baseAsset, strategy.config.Bridge, ratio)
	}

	return ratios, nil
}

// Jump from the source coin to the destination coin through bridge coin
func (strategy *testStrategy) jumpBestCoin(bestRatio *ratioPair) error {
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
