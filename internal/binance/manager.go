package binance

import (
	"binance-trade-bot/internal"
	binanceSpot "github.com/dirname/binance/spot/client"
	binanceSpotOrderSide "github.com/dirname/binance/spot/client/orderSide"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"math"
	"strconv"
	"strings"
)

type Manager interface {
	GetAccount() (*binanceSpot.AccountInfoResponse, error)
	GetTickerPrices() (map[string]float64, error)
	GetTickerPrice(symbol string) (float64, error)
	GetOpenOrder(symbol string) (float64, error)
	GetTradeFees() (map[string]decimal.Decimal, error)
	GetFee(originSymbol string, targetSymbol string, currentBalance float64, targetBalance float64, baseFee decimal.Decimal, usingBnbForFees bool, optionalCoinPrice float64, selling bool) (float64, error)
	GetUsingBnbForFees() (bool, error)
	GetCurrencyBalance(symbol string /*, forceNotCache bool*/) (float64, error)
	GetAltTick(originSymbol string, targetSymbol string) (int, error)
	GetMinNotional(originSymbol string, targetSymbol string) (float64, error)
	GetSymbolFilter(originSymbol string, targetSymbol string, filterType string) (map[string]interface{}, error)
	SellQuantity(originSymbol string, targetSymbol string, originBalance float64) (float64, error)
	BuyQuantity(originSymbol string, targetSymbol string, targetBalance float64, fromCoinPrice float64) (float64, error)
	BuyAlt(originSymbol string, targetSymbol string) (*binanceSpot.NewOrderResponseResult, error)
	SellAlt(originSymbol string, targetSymbol string) (*binanceSpot.NewOrderResponseResult, error)

	GetCurrentCoin() (string, error)
	GetSupportedSymbols() (map[string]map[string]*ExchangeSymbol, error)
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

	var cache cacheClient
	if config.UseRealtime {
		cache, err = newCacheClient(config, chanStop)
		if err != nil {
			logrus.Fatal("Couldn't access Socket API -  [error: %v]", err)
		}
	}

	return ManagerImpl{
		api,
		cache,
	}, nil
}

func (manager ManagerImpl) GetAccount() (*binanceSpot.AccountInfoResponse, error) {
	resp, err := manager.api.getAccountInfo()
	if err != nil {
		logrus.Errorf("GetAccount %v", err)
		return nil, err
	}

	return resp, nil
}

func (manager ManagerImpl) GetTickerPrices() (map[string]float64, error) {
	resp, err := manager.api.getSymbolTickerPrice("")
	if err != nil {
		logrus.Errorf("GetSymbolTickerPrice %v", err)
		return nil, err
	}

	prices := make(map[string]float64)
	for _, item := range resp {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return nil, err
		}

		prices[item.Symbol] = price
	}
	return prices, err
}

func (manager ManagerImpl) GetTickerPrice(symbol string) (float64, error) {
	resp, err := manager.api.getSymbolTickerPrice(symbol)
	if err != nil {
		logrus.Errorf("GetSymbolTickerPrice %v", err)
		return 0, err
	}

	var symbolPrice *binanceSpot.SymbolPriceTickerResponse
	for key := range resp {
		if resp[key].Symbol == symbol {
			symbolPrice = &resp[key]
			break
		}
	}
	if symbolPrice == nil {
		err = errors.Errorf("GetSymbolTickerPrice symbol %s not found", symbol)
		return 0, err
	}

	return strconv.ParseFloat(symbolPrice.Price, 64)
}

func (manager ManagerImpl) GetOpenOrder(symbol string) (float64, error) {
	_, err := manager.api.getOpenOrder(symbol)
	if err != nil {
		logrus.Errorf("GetSymbolTickerPrice %v", err)
		return 0, err
	}

		err = errors.Errorf("GetSymbolTickerPrice symbol %s not found", symbol)
		return 0, err
}


func (manager ManagerImpl) GetTradeFees() (map[string]decimal.Decimal, error) {
	resp, err := manager.api.sapiTradeFee("")
	if err != nil {
		logrus.Errorf("GetAccount %v", err)
		return nil, err
	}

	var fees = make(map[string]decimal.Decimal)
	for key := range resp {
		fee := resp[key]
		// https://github.com/edeng23/binance-trade-bot/blob/d76e4da7b5b41a925cf06d3c33c34b26db191932/binance_trade_bot/binance_api_manager.py#L43
		fees[fee.Symbol] = fee.TakerCommission
	}

	return fees, nil
}

func (manager ManagerImpl) GetFee(originSymbol string, targetSymbol string, currentBalance float64, targetBalance float64, baseFee decimal.Decimal, usingBnbForFees bool, optionalCoinPrice float64, selling bool) (float64, error) {
	if !usingBnbForFees {
		baseFeeFloat64, _ := baseFee.Float64()
		return baseFeeFloat64, nil
	}

	var amountTrading float64
	var err error
	if selling {
		amountTrading, err = manager.SellQuantity(originSymbol, targetSymbol, currentBalance)
	} else {
		amountTrading, err = manager.BuyQuantity(originSymbol, targetSymbol, targetBalance, optionalCoinPrice)
	}
	if err != nil {
		return 0, err
	}

	var feeAmountBNB decimal.Decimal
	// feeAmount := amountTrading * baseFee * 0.75
	feeAmount := baseFee.Mul(decimal.NewFromFloat(amountTrading)).Mul(decimal.NewFromFloat(0.75))
	if originSymbol == "BNB" {
		feeAmountBNB = feeAmount
	} else {
		originPrice, err := manager.GetTickerPrice(originSymbol + "BNB")
		if err != nil {
			return 0, err
		} else if originPrice == 0 {
			feeAmountFloat64, _ := feeAmount.Float64()
			logrus.Errorf("GetTickerPrice for %s is 0", originSymbol+"BNB")
			return feeAmountFloat64, nil
		}

		feeAmountBNB = feeAmount.Mul(decimal.NewFromFloat(originPrice))
	}

	bnbBalance, err := manager.GetCurrencyBalance("BNB")
	if err != nil {
		return 0, err
	}

	feeAmountBNBFloat64, _ := feeAmountBNB.Float64()
	if bnbBalance >= feeAmountBNBFloat64 {
		bnbBalance = bnbBalance * 0.75
	}
	return bnbBalance, nil
}

func (manager ManagerImpl) GetUsingBnbForFees() (bool, error) {
	resp, err := manager.api.getBnbBurn()
	if err != nil {
		logrus.Errorf("GetAccount %v", err)
		return false, err
	}

	return resp.SpotBNBBurn, nil
}

func (manager ManagerImpl) GetCurrencyBalance(symbol string /*, forceNotCache bool*/) (float64, error) {
	resp, err := manager.api.getAccountInfo()
	if err != nil {
		logrus.Errorf(" %v", err)
		return 0, err
	}

	var currencyBalance = make(map[string]string)
	for key := range resp.Balances {
		balance := resp.Balances[key]
		currencyBalance[balance.Asset] = balance.Free
	}

	if _, ok := currencyBalance[symbol]; !ok {
		err := errors.Errorf("Coin symbol %v not found", symbol)
		logrus.Error(err)
		return 0, err
	}

	return strconv.ParseFloat(currencyBalance[symbol], 64)
}

func (manager ManagerImpl) GetAltTick(originSymbol string, targetSymbol string) (int, error) {
	filter, err := manager.GetSymbolFilter(originSymbol, targetSymbol, "LOT_SIZE")
	if err != nil {
		return 0, err
	}

	stepSizeInterface, ok := filter["stepSize"]
	if !ok {
		err := errors.Errorf("GetAltTick stepSize is not defined")
		return 0, err
	}

	stepSize, ok := stepSizeInterface.(string)
	if !ok {
		err := errors.Errorf("stepSizeInterface can't convert to string")
		return 0, err
	}

	if 0 == strings.Index(stepSize, "1") {
		return 1 - strings.Index(stepSize, "."), nil
	}
	return strings.Index(stepSize, "1") - 1, nil
}

func (manager ManagerImpl) GetMinNotional(originSymbol string, targetSymbol string) (float64, error) {
	filter, err := manager.GetSymbolFilter(originSymbol, targetSymbol, "MIN_NOTIONAL")
	if err != nil {
		return 0, err
	}

	minNotionalInterface, ok := filter["minNotional"]
	if !ok {
		err := errors.Errorf("GetAltTick stepSize is not defined")
		return 0, err
	}

	minNotional, ok := minNotionalInterface.(float64)
	if !ok {
		err := errors.Errorf("minNotionalInterface can't convert to float64")
		return 0, err
	}

	return minNotional, nil
}

func (manager ManagerImpl) GetSymbolFilter(originSymbol string, targetSymbol string, filterType string) (map[string]interface{}, error) {
	resp, err := manager.api.getExchangeInfo()
	if err != nil {
		err = errors.Errorf("GetSymbolFilter error %v", err)
		return nil, err
	}

	var filters *[]map[string]interface{}
	symbolConversion := originSymbol + targetSymbol
	for key := range resp.Symbols {
		symbol := resp.Symbols[key]
		if symbol.Symbol == symbolConversion {
			filters = &symbol.Filters
		}
	}
	if filters == nil {
		err = errors.Errorf("GetSymbolFilter symbol %s not found in %v", symbolConversion, resp)
		return nil, err
	}

	var filterForType *map[string]interface{}
	for key := range *filters {
		filter := (*filters)[key]
		if filterItemForType, ok := filter["filterType"]; ok && filterType == filterItemForType {
			filterForType = &filter
		}
	}
	if filterForType == nil {
		err = errors.Errorf("GetSymbolFilter filter for type %s not found in %v %v", filterType, *filters, resp)
		return nil, err
	}

	return *filterForType, nil
}

func (manager ManagerImpl) SellQuantity(originSymbol string, targetSymbol string, originBalance float64) (float64, error) {
	if originBalance == 0 {
		err := errors.Errorf("Currency balance %s is empty", originSymbol)
		logrus.Error(err)
		return 0, err
	}

	originTick, err := manager.GetAltTick(originSymbol, targetSymbol)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return math.Floor(originBalance*10*float64(originTick)) / float64(10*originTick), nil
}

func (manager ManagerImpl) BuyQuantity(originSymbol string, targetSymbol string, targetBalance float64, fromCoinPrice float64) (float64, error) {
	if targetBalance == 0 {
		err := errors.Errorf("Currency balance %s is empty.", targetSymbol)
		logrus.Error(err)
		return 0, err
	}

	if fromCoinPrice == 0 {
		err := errors.Errorf("Coin price from %s to %s is zero.", originSymbol, targetSymbol)
		logrus.Error(err)
		return 0, err
	}

	originTick, err := manager.GetAltTick(originSymbol, targetSymbol)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return math.Floor(targetBalance*10*float64(originTick)/fromCoinPrice) / float64(10*originTick), nil
}

// BuyAlt - buy alt coin
func (manager ManagerImpl) BuyAlt(originSymbol string, targetSymbol string) (*binanceSpot.NewOrderResponseResult, error) {
	originBalance, err := manager.GetCurrencyBalance(originSymbol)
	if err != nil {
		return nil, err
	}

	targetBalance, err := manager.GetCurrencyBalance(targetSymbol)
	if err != nil {
		return nil, err
	}

	fromCoinPrice, err := manager.GetTickerPrice(originSymbol + targetSymbol)
	if err != nil {
		return nil, err
	}

	orderQuantity, err := manager.BuyQuantity(originSymbol, targetSymbol, originBalance, targetBalance)
	if err != nil {
		return nil, err
	}

	resp, err := manager.api.newOrderFull(originSymbol+targetSymbol, binanceSpotOrderSide.Buy, "", "", "",
		decimal.NewFromFloat(orderQuantity), decimal.Decimal{}, decimal.NewFromFloat(fromCoinPrice), decimal.Decimal{}, decimal.Decimal{})

	return resp, nil
}

// SellAlt - sell alt coin
func (manager ManagerImpl) SellAlt(originSymbol string, targetSymbol string) (*binanceSpot.NewOrderResponseResult, error) {
	originBalance, err := manager.GetCurrencyBalance(originSymbol)
	if err != nil {
		return nil, err
	}

	fromCoinPrice, err := manager.GetTickerPrice(originSymbol + targetSymbol)
	if err != nil {
		return nil, err
	}

	orderQuantity, err := manager.SellQuantity(originSymbol, targetSymbol, originBalance)
	if err != nil {
		return nil, err
	}

	resp, err := manager.api.newOrderFull(originSymbol+targetSymbol, binanceSpotOrderSide.Sell, "", "", "",
		decimal.NewFromFloat(orderQuantity), decimal.Decimal{}, decimal.NewFromFloat(fromCoinPrice), decimal.Decimal{}, decimal.Decimal{})

	return resp, nil
}

func (manager ManagerImpl) GetCurrentCoin() (string, error) {
	resp, err := manager.api.getAccountInfo()
	if err != nil {
		logrus.Errorf("getAccountInfo %v", err)
		return "", err
	}

	var symbolForMaxFree string
	var maxFree float64
	for key := range resp.Balances {
		balance := resp.Balances[key]
		currentMaxFree, err := strconv.ParseFloat(balance.Free, 64)
		if err != nil {
			logrus.Errorf("ParseFloat for fee in %v with error %v", balance, err)
			return "", err
		}

		if currentMaxFree >= maxFree {
			symbolForMaxFree = balance.Asset
			maxFree = currentMaxFree
		}
	}

	return symbolForMaxFree, nil
}

// GetSupportedSymbols - return all possible symbols
func (manager ManagerImpl) GetSupportedSymbols() (map[string]map[string]*ExchangeSymbol, error) {
	resp, err := manager.api.getExchangeInfo()
	if err != nil {
		err = errors.Errorf("GetSymbolFilter error %v", err)
		return nil, err
	}

	var supportedSymbols = make(map[string]map[string]*ExchangeSymbol)
	for _, symbol := range resp.Symbols {
		if symbol.Status != "TRADING" {
			continue
		}
		if !symbol.IsSpotTradingAllowed {
			continue
		}

		if supportedSymbols[symbol.BaseAsset] == nil {
			supportedSymbols[symbol.BaseAsset] = make(map[string]*ExchangeSymbol)
		}
		supportedSymbols[symbol.BaseAsset][symbol.QuoteAsset] = &ExchangeSymbol{
			BaseAsset:  symbol.BaseAsset,
			QuoteAsset: symbol.QuoteAsset,
			Filters:    symbol.Filters,
		}
	}

	return supportedSymbols, nil
}
