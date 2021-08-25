package strategies

import (
	"binance-trade-bot/internal"
	"binance-trade-bot/internal/binance"
)

type strategyImpl struct {
	binanceManager binance.Manager
	config        internal.AppConfig
}

func NewStrategy(binanceManager binance.Manager, config internal.AppConfig) *strategyImpl {
	return &strategyImpl{
		binanceManager,
		config,
	}
}
