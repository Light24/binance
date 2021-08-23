package main

import (
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

	config := internal.InitConfig()
	binanceClient := internal.NewBinanceClient(config)

	accountInfo, err := binanceClient.GetAccountInfo()
	if err != nil {
		logrus.Fatal("GetAccountInfo error: %v", err)
	}
	_ = accountInfo
}
