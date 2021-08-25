package binance

import (
	"binance-trade-bot/internal"
	"github.com/dirname/binance/model"
	binanceMarketWebsocket "github.com/dirname/binance/spot/websocket/market"
	spotclient "github.com/dirname/binance/spot/websocket/market"
	"github.com/sirupsen/logrus"
	"time"
)

type realtimeMarketClient interface {
	close()
}

type realtimeMarketClientImpl struct {
	*binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient
}

func newBinanceRealtimeMarketClient(config internal.AppConfig) (realtimeMarketClient, error) {
	client := realtimeMarketClientImpl{
		SpotAllMarketMiniTickerWebsocketClient: newSpotAllMarketMiniTickerWebsocketClient(config.HostWss, "!miniTicker@arr"),
	}

	client.SetReadTimerInterval(2 * time.Second)
	client.SetReconnectWaitTime(10 * time.Minute)
	client.SetKeepAliveInterval(10 * time.Second)
	client.SetHandler(func() {
		// TODO reset 24hrMiniTicker
	}, func(response interface{}) {
		switch response.(type) {
		case spotclient.AllMarketMiniTickerResponse:
			// TODO 24hrMiniTicker
			logrus.Info("AllMarketMiniTicker Response: %v", response.(spotclient.AllMarketMiniTickerResponse))
		case spotclient.AllMarketMiniTickerCombinedResponse:
			// TODO 24hrMiniTicker
			logrus.Info("AllMarketMiniTickerCombinedResponse: %v", response.(spotclient.AllMarketMiniTickerCombinedResponse))
		case model.WebsocketCommonResponse:
			logrus.Info("Websocket Response: %v", response.(model.WebsocketCommonResponse))
		case model.WebsocketErrorResponse:
			logrus.Info("Websocket Error Response: %v", response.(model.WebsocketErrorResponse))
		default:
			logrus.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)

	return client, nil
}

func newSpotAllMarketMiniTickerWebsocketClient(host string, streams ...string) *binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient {
	c := new(binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}

func (c realtimeMarketClientImpl) close() {
	c.Close()
	logrus.Info("Client closed")
}
