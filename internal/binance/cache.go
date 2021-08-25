package binance

import (
	"binance-trade-bot/internal"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type cacheClient interface {
}

type cacheClientImpl struct {
	realtimeAccountClient
	realtimeBalanceClient
	realtimeMarketClient
	realtimeOrderClient
}

func newCacheClient(config internal.AppConfig, stop chan interface{}) (cacheClient, error) {
	listenKeyBuilder := binanceAccountWebsocket.NewListenKeyBuilder(config.HostRest, config.ApiKey, config.ApiSecret)
	key, err := listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		err = errors.Errorf("CreateSpotListenKey Failed to create spot listen key: %s", err.Error())
		logrus.Error(err)
		return nil, err
	}

	accountClient, err := newBinanceRealtimeAccountClient(config, key)
	if err != nil {
		return nil, err
	}

	balanceClient, err := newBinanceRealtimeBalanceClient(config, key)
	if err != nil {
		return nil, err
	}

	marketClient, err := newBinanceRealtimeMarketClient(config)
	if err != nil {
		return nil, err
	}

	orderClient, err := newBinanceRealtimeOrderClient(config, key)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case <-stop:
				accountClient.close()
				balanceClient.close()
				marketClient.close()
				orderClient.close()
				logrus.Info("Client closed")
				return
			case <-time.After(10 * time.Minute):
				errInterface, err := listenKeyBuilder.PingSpotListenKey(key)
				if err != nil {
					logrus.Errorf("PingSpotListenKey error %v %v", errInterface, err)
				}
			}
		}
	}()

	return &cacheClientImpl{
		accountClient,
		balanceClient,
		marketClient,
		orderClient,
	}, nil
}
