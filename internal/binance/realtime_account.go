package binance

import (
	"binance-trade-bot/internal"
	"github.com/dirname/binance/model"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/sirupsen/logrus"
	"time"
)

type realtimeAccountClient interface {
	close()
}

type realtimeAccountClientImpl struct {
	*binanceAccountWebsocket.AccountWebsocketClient
}

func newBinanceRealtimeAccountClient(config internal.AppConfig, key string) (realtimeAccountClient, error) {
	client := realtimeAccountClientImpl{
		AccountWebsocketClient: newAccountWebsocketClient(config.HostWss, key),
	}

	client.SetReadTimerInterval(2 * time.Second)
	client.SetReconnectWaitTime(10 * time.Minute)
	client.SetKeepAliveInterval(10 * time.Second)
	client.SetHandler(func() {
		// TODO reset outboundAccountPosition
	}, func(response interface{}) {
		switch response.(type) {
		case binanceAccountWebsocket.AccountPosition:
			// TODO outboundAccountPosition
			logrus.Info("AccountUpdate Response: %v", response.(binanceAccountWebsocket.AccountPosition))
		case binanceAccountWebsocket.AccountCombinedPosition:
			// TODO outboundAccountPosition
			logrus.Info("AccountUpdateCombinedResponse: %v", response.(binanceAccountWebsocket.AccountCombinedPosition))
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

func newAccountWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.AccountWebsocketClient {
	c := new(binanceAccountWebsocket.AccountWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}

func (c realtimeAccountClientImpl) close() {
	c.Close()
	logrus.Info("Client closed")
}
