package binance

import (
	"binance-trade-bot/internal"
	binanceConfig "github.com/dirname/binance/config"
	binanceModel "github.com/dirname/binance/model"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type RealtimeOrderClient interface {
}

type RealtimeOrderClientImpl struct {
	*binanceAccountWebsocket.OrderWebsocketClient
	stop chan interface{}
}

func NewBinanceRealtimeOrderClient(config internal.AppConfig, stop chan interface{}) (RealtimeOrderClient, error) {
	hostRest, hostWss := binanceConfig.SpotRestHost, binanceConfig.SpotWssHost
	if config.TestNet {
		hostRest, hostWss = "testnet.binance.vision", "testnet.binance.vision"
	}

	listenKeyBuilder := binanceAccountWebsocket.NewListenKeyBuilder(hostRest, config.ApiKey, config.ApiSecret)
	key, err := listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		err = errors.Errorf("CreateSpotListenKey 1 Failed to create spot listen key: %s", err.Error())
		logrus.Error(err)
		return nil, err
	}

	key, err = listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		err = errors.Errorf("CreateSpotListenKey 2 Failed to create spot listen key: %s", err.Error())
		logrus.Error(err)
		return nil, err
	}

	client := RealtimeOrderClientImpl{
		OrderWebsocketClient: newOrderWebsocketClient(hostWss, key),
		stop:                 stop,
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
	// client.Subscribe(1222, key)

	go func() {
		for {
			select {
			case <-client.stop:
				client.Close()
				logrus.Info("Client closed")
			case <-time.After(10 * time.Minute):
				errInterface, err := listenKeyBuilder.PingSpotListenKey(key)
				if err != nil {
					logrus.Errorf("PingSpotListenKey error %v %v", errInterface, err)
				}
			}
		}
	}()

	return client, nil
}

func newOrderWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.OrderWebsocketClient {
	c := new(binanceAccountWebsocket.OrderWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}
