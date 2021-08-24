package strategies

import (
	"binance-trade-bot/internal"
	"binance-trade-bot/internal/binance"
)

type strategyImpl struct {
	binanceClient binance.BinanceClient
	config        internal.AppConfig
}

func NewStrategy(binanceClient binance.BinanceClient, config internal.AppConfig) *strategyImpl {
	return &strategyImpl{
		binanceClient: binanceClient,
		config: config,
	}
}
