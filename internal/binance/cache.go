package binance

import (
	"binance-trade-bot/internal"
	binanceAccountWebsocket "github.com/dirname/binance/spot/websocket/account"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type CacheClient interface {
}

type cacheClientImpl struct {
	RealtimeAccountClient
	RealtimeBalanceClient
	RealtimeMarketClient
	RealtimeOrderClient
}

func NewCacheClient(config internal.AppConfig, stop chan interface{}) (CacheClient, error) {
	listenKeyBuilder := binanceAccountWebsocket.NewListenKeyBuilder(config.HostRest, config.ApiKey, config.ApiSecret)
	key, err := listenKeyBuilder.CreateSpotListenKey()
	if err != nil {
		err = errors.Errorf("CreateSpotListenKey Failed to create spot listen key: %s", err.Error())
		logrus.Error(err)
		return nil, err
	}

	accountClient, err := NewBinanceRealtimeAccountClient(config, key)
	if err != nil {
		return nil, err
	}

	balanceClient, err := NewBinanceRealtimeBalanceClient(config, key)
	if err != nil {
		return nil, err
	}

	marketClient, err := NewBinanceRealtimeMarketClient(config)
	if err != nil {
		return nil, err
	}

	orderClient, err := NewBinanceRealtimeOrderClient(config, key)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case <-stop:
				accountClient.Close()
				balanceClient.Close()
				marketClient.Close()
				orderClient.Close()
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
