package binance

import (
	"binance-trade-bot/internal"
	binanceModel "github.com/dirname/binance/model"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/sirupsen/logrus"
	"time"
)

type realtimeBalanceClient interface {
	close()
}

type realtimeBalanceClientImpl struct {
	*binanceAccountWebsocket.BalanceWebsocketClient
	stop chan interface{}
}

func newBinanceRealtimeBalanceClient(config internal.AppConfig, key string) (realtimeBalanceClient, error) {
	client := realtimeBalanceClientImpl{
		BalanceWebsocketClient: newBalanceWebsocketClient(config.HostWss, key),
	}

	client.SetReadTimerInterval(2 * time.Second)
	client.SetReconnectWaitTime(10 * time.Minute)
	client.SetKeepAliveInterval(10 * time.Second)
	client.SetHandler(func() {
		// TODO reset balanceUpdate
	}, func(response interface{}) {
		switch response.(type) {
		case binanceAccountWebsocket.BalanceUpdate:
			// TODO balanceUpdate
			logrus.Info("BalanceUpdate Response: %v", response.(binanceAccountWebsocket.BalanceUpdate))
		case binanceAccountWebsocket.BalanceCombinedUpdate:
			// TODO balanceUpdate
			logrus.Info("BalanceUpdateCombinedResponse: %v", response.(binanceAccountWebsocket.BalanceCombinedUpdate))
		case binanceModel.WebsocketCommonResponse:
			logrus.Info("Websocket Response: %v", response.(binanceModel.WebsocketCommonResponse))
		case binanceModel.WebsocketErrorResponse:
			logrus.Info("Websocket Error Response: %v", response.(binanceModel.WebsocketErrorResponse))
		default:
			logrus.Info("Unknown Response: %v", response)
		}
	})
	client.Connect(true)

	return client, nil
}

func newBalanceWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.BalanceWebsocketClient {
	c := new(binanceAccountWebsocket.BalanceWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}

func (c realtimeBalanceClientImpl) close() {
	c.Close()
	logrus.Info("Client closed")
}
