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
	UseRealtime bool

	strategy        string
	Bridge          string
	ScoutMultiplier float64
}

func InitConfig(testNet bool) AppConfig {
	var appConfig AppConfig
	appConfig.RecvTimeout = 10 * time.Second
	appConfig.UseRealtime = false
	appConfig.strategy = "default"
	appConfig.Bridge = "USDT"
	appConfig.ScoutMultiplier = 5

	if testNet {
		appConfig.ApiKey = "lgtrEDSGBbFm3eEoQL6oCrCNXRaWzAGJL4jAdvZIrtqqApWaBMKfW7WDnObIhGAt"
		appConfig.ApiSecret = "qrw3iV9BynFEdwrMmmP4vkvusWeLCv1fC6O50veodZ5gBRZ4uKFBNoYeVn2v0Xpz"
		appConfig.HostRest = "testnet.binance.vision"
		appConfig.HostWss = "testnet.binance.vision"
	} else {
		appConfig.ApiKey = "feBpm5gMBeKXi5SZkMmB5HDsEuC019ZQ6w4rARqy5e56E9O8f8WJYUS5Dy86ZLUI"
		appConfig.ApiSecret = "V3Rspz4VQW4RnlUvZMMiFNGM4Wer4IDGHfsA8UoGLRAZrw6e93TjPkQg4fCBB4BQ"
		appConfig.HostRest = binanceConfig.SpotRestHost
		appConfig.HostWss = binanceConfig.SpotWssHost

	}

	return appConfig
}
