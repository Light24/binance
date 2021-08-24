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

type RealtimeBalanceClient interface {
}

type RealtimeBalanceClientImpl struct {
	*binanceAccountWebsocket.BalanceWebsocketClient
	stop chan interface{}
}

func NewBinanceRealtimeBalanceClient(config internal.AppConfig, stop chan interface{}) (RealtimeBalanceClient, error) {
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

	client := RealtimeBalanceClientImpl{
		BalanceWebsocketClient: newBalanceWebsocketClient(hostWss, key),
		stop:                   stop,
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

func newBalanceWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.BalanceWebsocketClient {
	c := new(binanceAccountWebsocket.BalanceWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}
