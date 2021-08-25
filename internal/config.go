package internal

import (
	binanceConfig "github.com/dirname/binance/config"
	"time"
)

type AppConfig struct {
	HostRest    string
	HostWss     string
	ApiKey      string
	ApiSecret   string
	RecvTimeout time.Duration

	strategy string
}

func InitConfig() AppConfig {
	var appConfig AppConfig
	appConfig.ApiKey = "lgtrEDSGBbFm3eEoQL6oCrCNXRaWzAGJL4jAdvZIrtqqApWaBMKfW7WDnObIhGAt"
	appConfig.ApiSecret = "qrw3iV9BynFEdwrMmmP4vkvusWeLCv1fC6O50veodZ5gBRZ4uKFBNoYeVn2v0Xpz"
	appConfig.RecvTimeout = 10 * time.Second

	testNet := true
	appConfig.HostRest, appConfig.HostWss = binanceConfig.SpotRestHost, binanceConfig.SpotWssHost
	if testNet {
		appConfig.HostRest, appConfig.HostWss = "testnet.binance.vision", "testnet.binance.vision"
	}

	appConfig.strategy = "default"

	return appConfig
}
