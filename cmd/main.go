package main

import (
	"binance-trade-bot/internal/binance"
	"binance-trade-bot/internal/strategies"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"binance-trade-bot/internal"
)

func init() {
	logFile := false
	if logFile {
		fileWriter := lumberjack.Logger{}
		logrus.SetOutput(&fileWriter)
	}

	logrus.SetLevel(logrus.TraceLevel)
}

func main() {
	logrus.Debugf("Starting..")
	defer logrus.Debugf("Shutdown.")

	chanStop := make(chan interface{})

	config := internal.InitConfig()
	binanceClient := binance.NewBinanceClient(config)
	binanceRealtimeClient := binance.NewBinanceRealtimeClient(config, chanStop)

	accountInfo, err := binanceClient.GetAccountInfo()
	if err != nil {
		logrus.Fatal("Couldn't access Binance API - API keys may be wrong or lack sufficient permissions [error: %v]", err)
	}
	_ = accountInfo

	_ = binanceRealtimeClient

	strategies.NewStrategy(binanceClient, config)
}
