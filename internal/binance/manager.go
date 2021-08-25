package binance

import (
	"binance-trade-bot/internal"
	"github.com/sirupsen/logrus"
)

type Manager interface {
}

type ManagerImpl struct {
	api   apiClient
	cache cacheClient
}

func New(config internal.AppConfig, chanStop chan interface{}) (Manager, error) {
	api := newApiClient(config)
	_, err := api.getAccountInfo()
	if err != nil {
		return nil, err
	}

	cache, err := newCacheClient(config, chanStop)
	if err != nil {
		logrus.Fatal("Couldn't access Socket API -  [error: %v]", err)
	}

	return ManagerImpl{
		api,
		cache,
	}, nil
}
