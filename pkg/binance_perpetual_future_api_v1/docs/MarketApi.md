# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**24hrTickerPriceChangeStatisticsMARKETDATA**](MarketApi.md#24hrTickerPriceChangeStatisticsMARKETDATA) | **Get** /fapi/v1/ticker/24hr | 24hr Ticker Price Change Statistics (MARKET_DATA)
[**CheckServerTime**](MarketApi.md#CheckServerTime) | **Get** /fapi/v1/time | Check Server Time
[**CompositeIndexSymbolInformation**](MarketApi.md#CompositeIndexSymbolInformation) | **Get** /fapi/v1/indexInfo | Composite Index Symbol Information
[**CompressedAggregateTradesListMARKETDATA**](MarketApi.md#CompressedAggregateTradesListMARKETDATA) | **Get** /fapi/v1/aggTrades | Compressed/Aggregate Trades List (MARKET_DATA)
[**ContinuousContractKlineCandlestickData**](MarketApi.md#ContinuousContractKlineCandlestickData) | **Get** /fapi/v1/continuousKlines | Continuous Contract Kline/Candlestick Data
[**ExchangeInformationMARKETDATA**](MarketApi.md#ExchangeInformationMARKETDATA) | **Get** /fapi/v1/exchangeInfo | Exchange Information (MARKET_DATA)
[**GetFundingRateHistoryMARKETDATA**](MarketApi.md#GetFundingRateHistoryMARKETDATA) | **Get** /fapi/v1/fundingRate | Get Funding Rate History (MARKET_DATA)
[**GetallLiquidationOrdersMARKETDATA**](MarketApi.md#GetallLiquidationOrdersMARKETDATA) | **Get** /fapi/v1/allForceOrders | Get all Liquidation Orders (MARKET_DATA)
[**HistoricalBLVTNAVKlineCandlestick**](MarketApi.md#HistoricalBLVTNAVKlineCandlestick) | **Get** /fapi/v1/lvtKlines | Historical BLVT NAV Kline/Candlestick
[**IndexPriceKlineCandlestickData**](MarketApi.md#IndexPriceKlineCandlestickData) | **Get** /fapi/v1/indexPriceKlines | Index Price Kline/Candlestick Data
[**KlineCandlestickDataMARKETDATA**](MarketApi.md#KlineCandlestickDataMARKETDATA) | **Get** /fapi/v1/klines | Kline/Candlestick Data (MARKET_DATA)
[**LongShortRatioMARKETDATA**](MarketApi.md#LongShortRatioMARKETDATA) | **Get** /futures/data/globalLongShortAccountRatio | Long/Short Ratio (MARKET_DATA)
[**MarkPriceKlineCandlestickData**](MarketApi.md#MarkPriceKlineCandlestickData) | **Get** /fapi/v1/markPriceKlines | Mark Price Kline/Candlestick Data
[**MarkPriceMARKETDATA**](MarketApi.md#MarkPriceMARKETDATA) | **Get** /fapi/v1/premiumIndex | Mark Price (MARKET_DATA)
[**OldTradesLookupMARKETDATA**](MarketApi.md#OldTradesLookupMARKETDATA) | **Get** /fapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**OpenInterestMARKETDATA**](MarketApi.md#OpenInterestMARKETDATA) | **Get** /fapi/v1/openInterest | Open Interest (MARKET_DATA)
[**OpenInterestStatisticsMARKETDATA**](MarketApi.md#OpenInterestStatisticsMARKETDATA) | **Get** /futures/data/openInterestHist | Open Interest Statistics (MARKET_DATA)
[**OrderBookMARKETDATA**](MarketApi.md#OrderBookMARKETDATA) | **Get** /fapi/v1/depth | Order Book (MARKET_DATA)
[**RecentTradesListMARKETDATA**](MarketApi.md#RecentTradesListMARKETDATA) | **Get** /fapi/v1/trades | Recent Trades List (MARKET_DATA)
[**SymbolOrderBookTickerMARKETDATA**](MarketApi.md#SymbolOrderBookTickerMARKETDATA) | **Get** /fapi/v1/ticker/bookTicker | Symbol Order Book Ticker (MARKET_DATA)
[**TakerBuySellVolumeMARKETDATA**](MarketApi.md#TakerBuySellVolumeMARKETDATA) | **Get** /futures/data/takerlongshortRatio | Taker Buy/Sell Volume (MARKET_DATA)
[**TestConnectivity**](MarketApi.md#TestConnectivity) | **Get** /fapi/v1/ping | Test Connectivity
[**TopTraderLongShortRatioAccountsMARKETDATA**](MarketApi.md#TopTraderLongShortRatioAccountsMARKETDATA) | **Get** /futures/data/topLongShortAccountRatio | Top Trader Long/Short Ratio (Accounts) (MARKET_DATA)
[**TopTraderLongShortRatioPositionsMARKETDATA**](MarketApi.md#TopTraderLongShortRatioPositionsMARKETDATA) | **Get** /futures/data/topLongShortPositionRatio | Top Trader Long/Short Ratio (Positions) (MARKET_DATA)

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

# **CompositeIndexSymbolInformation**
> CompositeIndexSymbolInformation(ctx, contentType, xMBXAPIKEY)
Composite Index Symbol Information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.

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

# **GetFundingRateHistoryMARKETDATA**
> GetFundingRateHistoryMARKETDATA(ctx, contentType)
Get Funding Rate History (MARKET_DATA)

Mark Price and Funding Rate  Weight: 1

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

# **GetallLiquidationOrdersMARKETDATA**
> GetallLiquidationOrdersMARKETDATA(ctx, symbol, startTime, endTime, limit, contentType)
Get all Liquidation Orders (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
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

# **HistoricalBLVTNAVKlineCandlestick**
> HistoricalBLVTNAVKlineCandlestick(ctx, symbol, interval, contentType, xMBXAPIKEY)
Historical BLVT NAV Kline/Candlestick

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**| token name | 
  **interval** | **string**|  | 
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

# **IndexPriceKlineCandlestickData**
> IndexPriceKlineCandlestickData(ctx, pair, interval, contentType)
Index Price Kline/Candlestick Data

Kline/candlestick bars for the index price of a pair. Klines are uniquely identified by their open time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pair** | **string**|  | 
  **interval** | **string**| 1m, 1h, 1d | 
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
  **interval** | **string**| 1m, 1h, 1d | 
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
> LongShortRatioMARKETDATA(ctx, symbol, period, contentType, xMBXAPIKEY)
Long/Short Ratio (MARKET_DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
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

# **MarkPriceKlineCandlestickData**
> MarkPriceKlineCandlestickData(ctx, symbol, interval, contentType)
Mark Price Kline/Candlestick Data

Kline/candlestick bars for the mark price of a symbol. Klines are uniquely identified by their open time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **interval** | **string**| 1m, 1h, 1d | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkPriceMARKETDATA**
> MarkPriceMARKETDATA(ctx, contentType)
Mark Price (MARKET_DATA)

Mark Price and Funding Rate  Weight: 1

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
> OpenInterestStatisticsMARKETDATA(ctx, symbol, period, contentType, xMBXAPIKEY)
Open Interest Statistics (MARKET_DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
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

# **OrderBookMARKETDATA**
> OrderBookMARKETDATA(ctx, symbol, contentType)
Order Book (MARKET_DATA)

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

# **TakerBuySellVolumeMARKETDATA**
> TakerBuySellVolumeMARKETDATA(ctx, symbol, period, contentType, xMBXAPIKEY)
Taker Buy/Sell Volume (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
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
> TopTraderLongShortRatioAccountsMARKETDATA(ctx, symbol, period, contentType, xMBXAPIKEY)
Top Trader Long/Short Ratio (Accounts) (MARKET_DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
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

# **TopTraderLongShortRatioPositionsMARKETDATA**
> TopTraderLongShortRatioPositionsMARKETDATA(ctx, symbol, period, contentType, xMBXAPIKEY)
Top Trader Long/Short Ratio (Positions) (MARKET_DATA)

Get present open interest of a specific symbol.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **period** | **string**| \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot; | 
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

