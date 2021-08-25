package main

import (
	"binance-trade-bot/internal/binance"
	"binance-trade-bot/internal/strategies"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"

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
	binanceManager, err := binance.New(config, chanStop)
	if err != nil {
		logrus.Fatal("Couldn't access Binance API - API keys may be wrong or lack sufficient permissions [error: %v]", err)
	}

	binanceManager.GetAltTick("BNB", "USDT")

	binanceManager.GetTradeFees()
	binanceManager.GetTickerPrice("")
	binanceManager.GetAccount()

	strategies.NewStrategy(binanceManager, config)
	time.Sleep(time.Hour)
}
