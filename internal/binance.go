package internal

import (
	"encoding/json"
	"github.com/dirname/binance"
	binanceLogger "github.com/dirname/binance/logging"
	"github.com/dirname/binance/model"
	binanceSpot "github.com/dirname/binance/spot/client"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"net/http"
	"time"
)

type BinanceClient interface {
	// GetAccountInfo -
	GetAccountInfo() (*binanceSpot.AccountInfoResponse, error)
	// SAPITradeFee - get_trade_fee
	SAPITradeFee(symbol string) (*binanceSpot.SAPITradeFeeResponse, error)
	// GetBnbBurn - get_bnb_burn_spot_margin
	GetBnbBurn() (*GetBnbBurnResponse, error)
	// GetSymbolTickerPrice -
	GetSymbolTickerPrice(symbol string) (*binanceSpot.SymbolPriceTickerResponse, error)
	// GetExchangeInfo -
	GetExchangeInfo() (*binanceSpot.ExchangeInfoResponse, error)
	// NewOrder -
	NewOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (interface{}, error)
	// NewOCO -
	NewOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType string, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty decimal.Decimal) (*binanceSpot.NewOCOOrderResponse, error)
	// DeleteOrder -
	DeleteOrder(symbol, origClientOrderID, newClientOrderID string, orderID int64) (*binanceSpot.DeleteOrderResponse, error)
}

type BinanceClientImpl struct {
	*binanceSpot.WalletClient
	*binanceSpot.TradeClient
	*binanceSpot.MarketClient
	recvTimeout time.Duration
}

func NewBinanceClient(config AppConfig) BinanceClient {
	return &BinanceClientImpl{
		WalletClient: binanceSpot.NewWalletClient(config.Host, config.ApiKey, config.ApiSecret),
		TradeClient:  binanceSpot.NewTradeClient(config.Host, config.ApiKey, config.ApiSecret),
		MarketClient: binanceSpot.NewMarketClient(config.Host, config.ApiKey),
		recvTimeout:  config.RecvTimeout,
	}
}

func (c *BinanceClientImpl) GetAccountInfo() (*binanceSpot.AccountInfoResponse, error) {
	interfaceResp, err := c.TradeClient.GetAccountInfo(c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.AccountInfoResponse)
	return &resp, nil
}

func (c *BinanceClientImpl) SAPITradeFee(symbol string) (*binanceSpot.SAPITradeFeeResponse, error) {
	interfaceResp, err := c.WalletClient.SAPITradeFee(symbol, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.SAPITradeFeeResponse)
	return &resp, nil
}

func (c *BinanceClientImpl) GetBnbBurn() (*GetBnbBurnResponse, error) {
	interfaceResp, err := c.GetBnbBurnImpl(c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(*GetBnbBurnResponse)
	return resp, nil
}

func (c *BinanceClientImpl) GetSymbolTickerPrice(symbol string) (*binanceSpot.SymbolPriceTickerResponse, error) {
	interfaceResp, err := c.MarketClient.GetSymbolTickerPrice(symbol)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.SymbolPriceTickerResponse)
	return &resp, nil
}

func (c *BinanceClientImpl) GetExchangeInfo() (*binanceSpot.ExchangeInfoResponse, error) {
	interfaceResp, err := c.MarketClient.GetExchangeInfo()
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.ExchangeInfoResponse)
	return &resp, nil
}

func (c *BinanceClientImpl) NewOrder(symbol, side, orderType, timeInForce, newClientOderID, newOrderRespType string, quantity, quoteOrderQTY, price, stopPrice, icebergQTY decimal.Decimal) (interface{}, error) {
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

func (c *BinanceClientImpl) NewOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType string, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty decimal.Decimal) (*binanceSpot.NewOCOOrderResponse, error) {
	interfaceResp, err := c.TradeClient.NewOCO(symbol, listClientOrderID, side, limitClientOrderId, stopClientOrderId, stopLimitTimeInForce, newOrderRespType, quantity, price, limitIcebergQty, stopPrice, stopLimitPrice, stopIcebergQty, c.recvTimeout)
	if err != nil {
		return nil, err
	} else if resp, ok := interfaceResp.(model.APIErrorResponse); ok {
		return nil, errors.Errorf("Error: %s with code %d", resp.Message, resp.Code)
	}

	resp := interfaceResp.(binanceSpot.NewOCOOrderResponse)
	return &resp, nil
}

func (c *BinanceClientImpl) DeleteOrder(symbol, origClientOrderID, newClientOrderID string, orderID int64) (*binanceSpot.DeleteOrderResponse, error) {
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

func (c *BinanceClientImpl) GetBnbBurnImpl(recv time.Duration) (interface{}, error) {
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
