package binance

import (
	"binance-trade-bot/internal"
	"encoding/json"
	"github.com/dirname/binance"
	binanceLogger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	binanceSpot "github.com/dirname/binance/spot/client"
	binanceSpotOrderRespType "github.com/dirname/binance/spot/client/orderRespType"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

type apiClient interface {
	// getAccountInfo -
	getAccountInfo() (*binanceSpot.AccountInfoResponse, error)
	// sapiTradeFee - get_trade_fee
	sapiTradeFee(symbol string) (binanceSpot.SAPITradeFeeResponse, error)
	// getBnbBurn - get_bnb_burn_spot_margin
	getBnbBurn() (*GetBnbBurnResponse, error)
	// getSymbolTickerPrice -
	getSymbolTickerPrice(symbol string) ([]binanceSpot.SymbolPriceTickerResponse, error)
	// getExchangeInfo -
	getExchangeInfo() (*binanceSpot.ExchangeInfoResponse, error)
	// newOrder -
	newOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (interface{}, error)
	// newOrder -
	newOrderFull(symbol, side, orderType, timeInForce, newClientOderID string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (*binanceSpot.NewOrderResponseResult, error)
	// newOCO -
	newOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType string, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty decimal.Decimal) (*binanceSpot.NewOCOOrderResponse, error)
	// deleteOrder -
	deleteOrder(symbol, origClientOrderID, newClientOrderID string, orderID int64) (*binanceSpot.DeleteOrderResponse, error)
}

type apiClientImpl struct {
	*binanceSpot.WalletClient
	*binanceSpot.TradeClient
	*binanceSpot.MarketClient
	recvTimeout time.Duration
}

func newApiClient(config internal.AppConfig) apiClient {
	return &apiClientImpl{
		WalletClient: binanceSpot.NewWalletClient(config.HostRest, config.ApiKey, config.ApiSecret),
		TradeClient:  binanceSpot.NewTradeClient(config.HostRest, config.ApiKey, config.ApiSecret),
		MarketClient: binanceSpot.NewMarketClient(config.HostRest, config.ApiKey),
		recvTimeout:  config.RecvTimeout,
	}
}

func (c *apiClientImpl) getAccountInfo() (*binanceSpot.AccountInfoResponse, error) {
	interfaceResp, err := c.TradeClient.GetAccountInfo(c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.AccountInfoResponse)
	return &resp, nil
}

func (c *apiClientImpl) sapiTradeFee(symbol string) (binanceSpot.SAPITradeFeeResponse, error) {
	interfaceResp, err := c.WalletClient.SAPITradeFee(symbol, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.SAPITradeFeeResponse)
	return resp, nil
}

func (c *apiClientImpl) getBnbBurn() (*GetBnbBurnResponse, error) {
	interfaceResp, err := c.getBnbBurnImpl(c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(*GetBnbBurnResponse)
	return resp, nil
}

func (c *apiClientImpl) getSymbolTickerPrice(symbol string) ([]binanceSpot.SymbolPriceTickerResponse, error) {
	interfaceResp, err := c.MarketClient.GetSymbolTickerPrice(symbol)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	var resp []binanceSpot.SymbolPriceTickerResponse
	if symbolPriceTicker, ok := interfaceResp.(binanceSpot.SymbolPriceTickerResponse); ok {
		resp = append(resp, symbolPriceTicker)
	} else {
		if symbolPriceTicker, ok := interfaceResp.([]binanceSpot.SymbolPriceTickerResponse); ok {
			resp = symbolPriceTicker
		}
	}

	return resp, nil
}

func (c *apiClientImpl) getExchangeInfo() (*binanceSpot.ExchangeInfoResponse, error) {
	interfaceResp, err := c.MarketClient.GetExchangeInfo()
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.ExchangeInfoResponse)
	return &resp, nil
}

func (c *apiClientImpl) newOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (interface{}, error) {
	interfaceResp, err := c.TradeClient.NewOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType, quantity, quoteOrderQTY, price, stopPrice, icebergQTY, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	if resp, ok := interfaceResp.(binanceSpot.NewOrderResponseResult); ok {
		return resp, nil
	} else if resp, ok := interfaceResp.(binanceSpot.NewOrderResponseFull); ok {
		return resp, nil
	} else if resp, ok := interfaceResp.(binanceSpot.NewOrderResponseACK); ok {
		return resp, nil
	}

	return nil, errors.Errorf("Imposibble situations: %v", interfaceResp)
}

func (c *apiClientImpl) newOrderFull(symbol, side, orderType, timeInForce, newClientOderID string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (*binanceSpot.NewOrderResponseResult, error) {
	respInterface, err := c.newOrder(symbol, side, orderType, timeInForce, newClientOderID, binanceSpotOrderRespType.Full, quantity, quoteOrderQTY, price, stopPrice, icebergQTY)
	if err != nil {
		return nil, err
	}

	resp, ok := respInterface.(binanceSpot.NewOrderResponseResult)
	if !ok {
		err := errors.Errorf("Unable get resp %v", resp)
		return nil, err
	}

	return &resp, nil
}

func (c *apiClientImpl) newOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType string, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty decimal.Decimal) (*binanceSpot.NewOCOOrderResponse, error) {
	interfaceResp, err := c.TradeClient.NewOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.NewOCOOrderResponse)
	return &resp, nil
}

func (c *apiClientImpl) deleteOrder(symbol, origClientOrderID, newClientOrderID string, orderID int64) (*binanceSpot.DeleteOrderResponse, error) {
	interfaceResp, err := c.TradeClient.DeleteOrder(symbol, origClientOrderID, newClientOrderID, orderID, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.DeleteOrderResponse)
	return &resp, nil
}

// extend binance

// GetBnbBurnResponse https://binance-docs.github.io/apidocs/spot/en/#get-bnb-burn-status-user_data
type GetBnbBurnResponse struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

func (c *apiClientImpl) getBnbBurnImpl(recv time.Duration) (interface{}, error) {
	var err error
	req, err := c.WalletClient.Builder.Build(http.MethodGet, "/sapi/v1/bnbBurn", "", true, true, recv)
	if err != nil {
		binanceLogger.Error("Failed to build url: %s", err.Error())
	}
	res, err := binance.HttpRequest(req)
	var parser map[string]interface{}
	err = json.Unmarshal(res, &parser)
	if _, ok := parser["code"]; ok {
		result := model.APIErrorResponse{}
		err = json.Unmarshal(res, &result)
		return result, err
	}

	result := GetBnbBurnResponse{}
	err = json.Unmarshal(res, &result)
	return result, err
}
