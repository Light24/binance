package internal

import "time"

type AppConfig struct {
	Host        string
	ApiKey      string
	ApiSecret   string
	RecvTimeout time.Duration
}

func InitConfig() AppConfig {
	var appConfig AppConfig
	appConfig.Host = "testnet.binance.vision"
	appConfig.ApiKey = "lgtrEDSGBbFm3eEoQL6oCrCNXRaWzAGJL4jAdvZIrtqqApWaBMKfW7WDnObIhGAt"
	appConfig.ApiSecret = "qrw3iV9BynFEdwrMmmP4vkvusWeLCv1fC6O50veodZ5gBRZ4uKFBNoYeVn2v0Xpz"
	appConfig.RecvTimeout = 5000 // 5 * time.Second

	return appConfig
}
