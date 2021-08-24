package binance

import (
	"binance-trade-bot/internal"
	binanceConfig "github.com/dirname/binance/config"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"time"
)

type BinanceRealtimeClient interface {
}

type binanceRealtimeClientImpl struct {
	*binanceAccountWebsocket.AccountWebsocketClient
}

func NewBinanceRealtimeClient(config internal.AppConfig, stop chan interface{}) BinanceRealtimeClient {
	host := binanceConfig.SpotWssHost
	if config.TestNet {
		host = "testnet.binance.com"
	}

	NewBinanceRealtimeBalanceClient(config, stop)
	//NewBinanceRealtimeOrderClient(config, stop)
	//NewBinanceRealtimeAccountClient(config, stop)
	//NewBinanceRealtimeMarketClient(config, stop)
	time.Sleep(time.Hour)

	listenKeyBuilder := binanceAccountWebsocket.NewListenKeyBuilder(host, config.ApiKey, config.ApiSecret)
	listenKeyBuilder.CreateSpotListenKey()

	return &binanceRealtimeClientImpl{
		// AccountWebsocketClient: newAccountWebsocketClient(host, "!userData@arr", "!miniTicker@arr"),
	}
}
