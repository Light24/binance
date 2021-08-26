package strategies

import (
	"binance-trade-bot/internal"
	"binance-trade-bot/internal/binance"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Strategy interface {
	Scout() error
}

type strategyImpl struct {
	binanceManager binance.Manager
	config         internal.AppConfig
}

type ratioPair struct {
	origin string
	target string
	ratio  float64
}

func NewStrategy(binanceManager binance.Manager, config internal.AppConfig) Strategy {
	return &strategyImpl{
		binanceManager,
		config,
	}
}

// Scout - Scout for potential jumps from the current coin to another coin
func (strategy *strategyImpl) Scout() error {
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

func (strategy *strategyImpl) jumpToBestCoin(currentSymbol string, currentCoinPrice float64) error {
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
func (strategy *strategyImpl) getRatios(currentSymbol string, currentCoinPrice float64) ([]ratioPair, error) {
	supportedSymbols, err := strategy.binanceManager.GetSupportedSymbols()
	if err != nil {
		return nil, err
	}

	var ratios []ratioPair
	for targetSymbol := range supportedSymbols {
		if currentSymbol == targetSymbol {
			continue
		} else if targetSymbol == strategy.config.Bridge {
			continue
		}

		optionalCoinPrice, err := strategy.binanceManager.GetTickerPrice(targetSymbol + strategy.config.Bridge)
		if err != nil {
			continue
		} else if optionalCoinPrice == 0 {
			continue
		}

		transactionFee, err := strategy.binanceManager.GetFee(currentSymbol, strategy.config.Bridge, false)
		if err != nil {
			continue
		}

		// https://github.com/edeng23/binance-trade-bot/blob/d76e4da7b5b41a925cf06d3c33c34b26db191932/binance_trade_bot/auto_trader.py#L131
		// https://github.com/edeng23/binance-trade-bot/blob/d76e4da7b5b41a925cf06d3c33c34b26db191932/binance_trade_bot/auto_trader.py#L71
		coinOptCoinRatio := currentCoinPrice / optionalCoinPrice
		ratio := coinOptCoinRatio*(1-transactionFee*strategy.config.ScoutMultiplier) - (currentCoinPrice / optionalCoinPrice)
		ratios = append(ratios, ratioPair{currentSymbol, targetSymbol, ratio})
	}

	return ratios, nil
}

// Jump from the source coin to the destination coin through bridge coin
func (strategy *strategyImpl) transactionThroughBridge(bestRatio *ratioPair) error {
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

	canSell := balance * fromCoinPrice > minNotional
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
