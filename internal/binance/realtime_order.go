package binance

import (
	"binance-trade-bot/internal"
	binanceModel "github.com/dirname/binance/model"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/sirupsen/logrus"
	"time"
)

type realtimeOrderClient interface {
	close()
}

type realtimeOrderClientImpl struct {
	*binanceAccountWebsocket.OrderWebsocketClient
}

func newBinanceRealtimeOrderClient(config internal.AppConfig, key string) (realtimeOrderClient, error) {
	client := realtimeOrderClientImpl{
		OrderWebsocketClient: newOrderWebsocketClient(config.HostWss, key),
	}

	client.SetReadTimerInterval(2 * time.Second)
	client.SetReconnectWaitTime(10 * time.Minute)
	client.SetKeepAliveInterval(10 * time.Second)
	client.SetHandler(func() {
	}, func(response interface{}) {
		switch response.(type) {
		case binanceAccountWebsocket.ListStatus:
			logrus.Info("ListUpdate Response: %v", response.(binanceAccountWebsocket.ListStatus))
		case binanceAccountWebsocket.ListCombinedStatus:
			logrus.Info("ListCombinedUpdate Response: %v", response.(binanceAccountWebsocket.ListCombinedStatus))
		case binanceAccountWebsocket.ExecutionReport:
			// TODO executionReport
			logrus.Info("ExecutionUpdate Response: %v", response.(binanceAccountWebsocket.ExecutionReport))
		case binanceAccountWebsocket.ExecutionCombinedReport:
			// TODO executionReport
			logrus.Info("ExecutionUpdateCombinedResponse: %v", response.(binanceAccountWebsocket.ExecutionCombinedReport))
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

func newOrderWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.OrderWebsocketClient {
	c := new(binanceAccountWebsocket.OrderWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}

func (c realtimeOrderClientImpl) close() {
	c.Close()
	logrus.Info("Client closed")
}
