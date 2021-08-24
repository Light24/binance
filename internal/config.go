package internal

import "time"

type AppConfig struct {
	TestNet     bool
	ApiKey      string
	ApiSecret   string
	RecvTimeout time.Duration

	strategy string
}

func InitConfig() AppConfig {
	var appConfig AppConfig
	appConfig.TestNet = true
	appConfig.ApiKey = "lgtrEDSGBbFm3eEoQL6oCrCNXRaWzAGJL4jAdvZIrtqqApWaBMKfW7WDnObIhGAt"
	appConfig.ApiSecret = "qrw3iV9BynFEdwrMmmP4vkvusWeLCv1fC6O50veodZ5gBRZ4uKFBNoYeVn2v0Xpz"
	appConfig.RecvTimeout = 5000 // 5 * time.Second

	appConfig.strategy = "default"

	return appConfig
}
