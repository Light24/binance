# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**24hrTickerPriceChangeStatisticsMARKETDATA**](MarketApi.md#24hrTickerPriceChangeStatisticsMARKETDATA) | **Get** /dapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics (MARKET_DATA)
[**BasisMARKETDATA**](MarketApi.md#BasisMARKETDATA) | **Get** /futures/data/basis | Basis (MARKET DATA)
[**CheckServerTime**](MarketApi.md#CheckServerTime) | **Get** /dapi/v1/time | Check Server Time
[**CompressedAggregateTradesListMARKETDATA**](MarketApi.md#CompressedAggregateTradesListMARKETDATA) | **Get** /dapi/v1/aggTrades | Compressed/Aggregate Trades List (MARKET_DATA)
[**ContinuousContractKlineCandlestickData**](MarketApi.md#ContinuousContractKlineCandlestickData) | **Get** /dapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**ExchangeInformationMARKETDATA**](MarketApi.md#ExchangeInformationMARKETDATA) | **Get** /dapi/v1/exchangeInfo | Exchange Information (MARKET_DATA)
[**GetFundingRateHistoryofPerpetualFutures**](MarketApi.md#GetFundingRateHistoryofPerpetualFutures) | **Get** /dapi/v1/fundingRate | Get Funding Rate History of Perpetual Futures
[**GetallLiquidationOrdersMARKETDATA**](MarketApi.md#GetallLiquidationOrdersMARKETDATA) | **Get** /dapi/v1/allForceOrders | Get all Liquidation Orders (MARKET_DATA)
[**IndexPriceKlineCandlestickData**](MarketApi.md#IndexPriceKlineCandlestickData) | **Get** /dapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**IndexPriceandMarkPrice**](MarketApi.md#IndexPriceandMarkPrice) | **Get** /dapi/v1/premiumIndex | Index Price and Mark Price
[**KlineCandlestickDataMARKETDATA**](MarketApi.md#KlineCandlestickDataMARKETDATA) | **Get** /dapi/v1/klines | Kline/Candlestick Data (MARKET_DATA)
[**LongShortRatioMARKETDATA**](MarketApi.md#LongShortRatioMARKETDATA) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio (MARKET DATA)
[**MarkPriceKlineCandlestickData**](MarketApi.md#MarkPriceKlineCandlestickData) | **Get** /dapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**OldTradesLookupMARKETDATA**](MarketApi.md#OldTradesLookupMARKETDATA) | **Get** /dapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OpenInterestMARKETDATA**](MarketApi.md#OpenInterestMARKETDATA) | **Get** /dapi/v1/openInterest | Open Interest (MARKET_DATA)
[**OpenInterestStatisticsMARKETDATA**](MarketApi.md#OpenInterestStatisticsMARKETDATA) | **Get** /futures/data/openInterestHist | Open Interest Statistics (MARKET DATA)
[**OrderBookMARKETDATA**](MarketApi.md#OrderBookMARKETDATA) | **Get** /dapi/v1/depth | Order Book (MARKET_DATA)
[**RecentTradesListMARKETDATA**](MarketApi.md#RecentTradesListMARKETDATA) | **Get** /dapi/v1/trades | Recent Trades List (MARKET_DATA)
[**SymbolOrderBookTickerMARKETDATA**](MarketApi.md#SymbolOrderBookTickerMARKETDATA) | **Get** /dapi/v1/ticker/bookTicker | Symbol Order Book Ticker (MARKET_DATA)
[**SymbolPriceTickerMARKETDATA**](MarketApi.md#SymbolPriceTickerMARKETDATA) | **Get** /dapi/v1/ticker/price | Symbol Price Ticker (MARKET_DATA)
[**TakerBuySellVolumeMARKETDATA**](MarketApi.md#TakerBuySellVolumeMARKETDATA) | **Get** /futures/data/takerBuySellVol | Taker Buy/Sell Volume (MARKET DATA)
[**TestConnectivity**](MarketApi.md#TestConnectivity) | **Get** /dapi/v1/ping | Test Connectivity
[**TopTraderLongShortRatioAccountsMARKETDATA**](MarketApi.md#TopTraderLongShortRatioAccountsMARKETDATA) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts) (MARKET DATA)
[**TopTraderLongShortRatioPositionsMARKETDATA**](MarketApi.md#TopTraderLongShortRatioPositionsMARKETDATA) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions) (MARKET DATA)

# **24hrTickerPriceChangeStatisticsMARKETDATA**
> 24hrTickerPriceChangeStatisticsMARKETDATA(ctx, contentType)
24hr Ticker Price Change Statistics (MARKET_DATA)

24 hour rolling window price change statistics.  Careful when accessing this with no symbol.  Weight:  1 for a single symbol;  40 when the symbol parameter is omitted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BasisMARKETDATA**
> BasisMARKETDATA(ctx, pair, contractType, period, contentType)
Basis (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **contractType** | **string**| CURRENT_QUARTER,NEXT_QUARTER | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckServerTime**
> CheckServerTime(ctx, contentType)
Check Server Time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompressedAggregateTradesListMARKETDATA**
> CompressedAggregateTradesListMARKETDATA(ctx, symbol, contentType)
Compressed/Aggregate Trades List (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContinuousContractKlineCandlestickData**
> ContinuousContractKlineCandlestickData(ctx, pair, contractType, interval, contentType)
Continuous Contract Kline/Candlestick Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **contractType** | **string**|  | 
  **interval** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExchangeInformationMARKETDATA**
> ExchangeInformationMARKETDATA(ctx, contentType)
Exchange Information (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFundingRateHistoryofPerpetualFutures**
> GetFundingRateHistoryofPerpetualFutures(ctx, symbol, contentType)
Get Funding Rate History of Perpetual Futures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetallLiquidationOrdersMARKETDATA**
> GetallLiquidationOrdersMARKETDATA(ctx, symbol, pair, startTime, endTime, limit, contentType)
Get all Liquidation Orders (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **pair** | **string**|  | 
  **startTime** | **string**|  | 
  **endTime** | **string**|  | 
  **limit** | **int32**| Default:100 Max:1000 | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexPriceKlineCandlestickData**
> IndexPriceKlineCandlestickData(ctx, pair, interval, contentType)
Index Price Kline/Candlestick Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **interval** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexPriceandMarkPrice**
> IndexPriceandMarkPrice(ctx, contentType)
Index Price and Mark Price

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KlineCandlestickDataMARKETDATA**
> KlineCandlestickDataMARKETDATA(ctx, symbol, interval, contentType)
Kline/Candlestick Data (MARKET_DATA)

Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **interval** | **string**| 1m, 1h, 1d, etc | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LongShortRatioMARKETDATA**
> LongShortRatioMARKETDATA(ctx, pair, period, contentType)
Long/Short Ratio (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkPriceKlineCandlestickData**
> MarkPriceKlineCandlestickData(ctx, symbol, interval, contentType)
Mark Price Kline/Candlestick Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **interval** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OldTradesLookupMARKETDATA**
> OldTradesLookupMARKETDATA(ctx, symbol, contentType, xMBXAPIKEY)
Old Trades Lookup (MARKET_DATA)

This endpoint need your API key only, not the secret key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 
  **xMBXAPIKEY** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenInterestMARKETDATA**
> OpenInterestMARKETDATA(ctx, symbol, contentType)
Open Interest (MARKET_DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenInterestStatisticsMARKETDATA**
> OpenInterestStatisticsMARKETDATA(ctx, pair, contractType, period, contentType)
Open Interest Statistics (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **contractType** | **string**| ALL,CURRENT_QUARTER,NEXT_QUARTER | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderBookMARKETDATA**
> OrderBookMARKETDATA(ctx, symbol, limit, contentType)
Order Book (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **limit** | **int32**| Default 100; max 1000. Valid limits:[5, 10, 20, 50, 100, 500, 1000] | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecentTradesListMARKETDATA**
> RecentTradesListMARKETDATA(ctx, symbol, contentType)
Recent Trades List (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SymbolOrderBookTickerMARKETDATA**
> SymbolOrderBookTickerMARKETDATA(ctx, contentType)
Symbol Order Book Ticker (MARKET_DATA)

Best price/qty on the order book for a symbol or symbols.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SymbolPriceTickerMARKETDATA**
> SymbolPriceTickerMARKETDATA(ctx, contentType)
Symbol Price Ticker (MARKET_DATA)

Latest price for a symbol or symbols.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TakerBuySellVolumeMARKETDATA**
> TakerBuySellVolumeMARKETDATA(ctx, pair, contractType, period, contentType)
Taker Buy/Sell Volume (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **contractType** | **string**| ALL,CURRENT_QUARTER,NEXT_QUARTER | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestConnectivity**
> TestConnectivity(ctx, contentType)
Test Connectivity

Test connectivity to the Rest API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopTraderLongShortRatioAccountsMARKETDATA**
> TopTraderLongShortRatioAccountsMARKETDATA(ctx, pair, period, contentType)
Top Trader Long/Short Ratio (Accounts) (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopTraderLongShortRatioPositionsMARKETDATA**
> TopTraderLongShortRatioPositionsMARKETDATA(ctx, pair, period, contentType)
Top Trader Long/Short Ratio (Positions) (MARKET DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

