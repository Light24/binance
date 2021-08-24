package binance

import (
	"binance-trade-bot/internal"
	binanceConfig "github.com/dirname/binance/config"
	"github.com/dirname/binance/model"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type RealtimeAccountClient interface {
}

type RealtimeAccountClientImpl struct {
	*binanceAccountWebsocket.AccountWebsocketClient
	stop chan interface{}
}

func NewBinanceRealtimeAccountClient(config internal.AppConfig, stop chan interface{}) (RealtimeAccountClient, error) {
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

	client := RealtimeAccountClientImpl{
		AccountWebsocketClient: newAccountWebsocketClient(hostWss, key),
		stop:                   stop,
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
	// client.Subscribe(1222, key)

	// listenKeyBuilder.PingSpotListenKey(key)
	// client.SetPingHandler(nil)
	// client.SetPongHandler(nil)

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

func newAccountWebsocketClient(host string, streams ...string) *binanceAccountWebsocket.AccountWebsocketClient {
	c := new(binanceAccountWebsocket.AccountWebsocketClient)
	c.WebsocketClient.Init(host, streams...)
	return c
}
