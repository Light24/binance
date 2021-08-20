# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**24hrTickerPriceChangeStatistics**](MarketApi.md#24hrTickerPriceChangeStatistics) | **Get** /api/v3/ticker/24hr | 24hr Ticker Price Change Statistics
[**CheckServerTime**](MarketApi.md#CheckServerTime) | **Get** /api/v3/time | Check Server Time
[**CompressedAggregateTradesList**](MarketApi.md#CompressedAggregateTradesList) | **Get** /api/v3/aggTrades | Compressed/Aggregate Trades List
[**CurrentAveragePrice**](MarketApi.md#CurrentAveragePrice) | **Get** /api/v3/avgPrice | Current Average Price
[**ExchangeInformation**](MarketApi.md#ExchangeInformation) | **Get** /api/v3/exchangeInfo | Exchange Information
[**KlineCandlestickData**](MarketApi.md#KlineCandlestickData) | **Get** /api/v3/klines | Kline/Candlestick Data
[**OldTradeLookup**](MarketApi.md#OldTradeLookup) | **Get** /api/v3/historicalTrades | Old Trade Lookup
[**OrderBook**](MarketApi.md#OrderBook) | **Get** /api/v3/depth | Order Book
[**RecentTradesList**](MarketApi.md#RecentTradesList) | **Get** /api/v3/trades | Recent Trades List
[**SymbolOrderBookTicker**](MarketApi.md#SymbolOrderBookTicker) | **Get** /api/v3/ticker/bookTicker | Symbol Order Book Ticker
[**SymbolPriceTicker**](MarketApi.md#SymbolPriceTicker) | **Get** /api/v3/ticker/price | Symbol Price Ticker
[**TestConnectivity**](MarketApi.md#TestConnectivity) | **Get** /api/v3/ping | Test Connectivity

# **24hrTickerPriceChangeStatistics**
> 24hrTickerPriceChangeStatistics(ctx, contentType)
24hr Ticker Price Change Statistics

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

Test connectivity to the Rest API and get the current server time.

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

# **CompressedAggregateTradesList**
> CompressedAggregateTradesList(ctx, symbol, contentType)
Compressed/Aggregate Trades List

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

# **CurrentAveragePrice**
> CurrentAveragePrice(ctx, symbol, contentType)
Current Average Price

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

# **ExchangeInformation**
> ExchangeInformation(ctx, contentType)
Exchange Information

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

# **KlineCandlestickData**
> KlineCandlestickData(ctx, symbol, interval, contentType)
Kline/Candlestick Data

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

# **OldTradeLookup**
> OldTradeLookup(ctx, symbol, contentType, xMBXAPIKEY)
Old Trade Lookup

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

# **OrderBook**
> OrderBook(ctx, symbol, contentType)
Order Book

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

# **RecentTradesList**
> RecentTradesList(ctx, symbol, contentType)
Recent Trades List

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

# **SymbolOrderBookTicker**
> SymbolOrderBookTicker(ctx, contentType)
Symbol Order Book Ticker

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

# **SymbolPriceTicker**
> SymbolPriceTicker(ctx, contentType)
Symbol Price Ticker

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

