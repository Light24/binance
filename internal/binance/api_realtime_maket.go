package binance

import (
	"binance-trade-bot/internal"
	binanceConfig "github.com/dirname/binance/config"
	"github.com/dirname/binance/model"
	binanceMarketWebsocket "github.com/dirname/binance/spot/websocket/market"
	spotclient "github.com/dirname/binance/spot/websocket/market"
	"github.com/sirupsen/logrus"
	"time"
)

type RealtimeMarketClient interface {
}

type RealtimeMarketClientImpl struct {
	*binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient
	stop chan interface{}
}

func NewBinanceRealtimeMarketClient(config internal.AppConfig, stop chan interface{}) (RealtimeMarketClient, error) {
	hostWss := binanceConfig.SpotRestHost
	if config.TestNet {
		hostWss = "testnet.binance.vision"
	}

	client := RealtimeMarketClientImpl{
		SpotAllMarketMiniTickerWebsocketClient: newSpotAllMarketMiniTickerWebsocketClient(hostWss, "!miniTicker@arr"),
		stop:                                   stop,
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
	// client.SetPingHandler(nil)
	// client.SetPongHandler(nil)

	go func() {
		select {
		case <-client.stop:
			//client.Unsubscribe(2, "!miniTicker@arr")
			client.Close()
			logrus.Info("Client closed")
		}
	}()

	return client, nil
}

func newSpotAllMarketMiniTickerWebsocketClient(host string, streams ...string) *binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient {
	c := new(binanceMarketWebsocket.SpotAllMarketMiniTickerWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}
