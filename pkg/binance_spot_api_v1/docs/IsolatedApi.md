# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewIsolatedMarginAccountListenKey**](IsolatedApi.md#CreateNewIsolatedMarginAccountListenKey) | **Post** /sapi/v1/userDataStream/isolated | Create New Isolated Margin Account Listen Key
[**DeleteIsolatedMarginAccountListenKey**](IsolatedApi.md#DeleteIsolatedMarginAccountListenKey) | **Delete** /sapi/v1/userDataStream/isolated | Delete Isolated Margin Account Listen Key
[**GetAllIsolatedMarginPairsMARKETDATA**](IsolatedApi.md#GetAllIsolatedMarginPairsMARKETDATA) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Pairs (MARKET_DATA)
[**GetIsolatedMarginAccountTransferHistoryUSERDATA**](IsolatedApi.md#GetIsolatedMarginAccountTransferHistoryUSERDATA) | **Get** /sapi/v1/margin/isolated/transfer | Get Isolated Margin Account Transfer History (USER_DATA)
[**IsolatedMarginAccountTransferMARGIN**](IsolatedApi.md#IsolatedMarginAccountTransferMARGIN) | **Post** /sapi/v1/margin/isolated/transfer | Isolated Margin Account Transfer (MARGIN)
[**QueryIsolatedMarginAccountInfoUSERDATA**](IsolatedApi.md#QueryIsolatedMarginAccountInfoUSERDATA) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**QueryIsolatedMarginPairMARKETDATA**](IsolatedApi.md#QueryIsolatedMarginPairMARKETDATA) | **Get** /sapi/v1/margin/isolated/pair | Query Isolated Margin Pair (MARKET_DATA)
[**RenewIsolatedMarginAccountListenKey**](IsolatedApi.md#RenewIsolatedMarginAccountListenKey) | **Put** /sapi/v1/userDataStream/isolated | Renew Isolated Margin Account Listen Key

# **CreateNewIsolatedMarginAccountListenKey**
> CreateNewIsolatedMarginAccountListenKey(ctx, symbol, contentType, xMBXAPIKEY)
Create New Isolated Margin Account Listen Key

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

# **DeleteIsolatedMarginAccountListenKey**
> DeleteIsolatedMarginAccountListenKey(ctx, listenKey, symbol, contentType, xMBXAPIKEY)
Delete Isolated Margin Account Listen Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenKey** | **string**|  | 
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

# **GetAllIsolatedMarginPairsMARKETDATA**
> GetAllIsolatedMarginPairsMARKETDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get All Isolated Margin Pairs (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
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

# **GetIsolatedMarginAccountTransferHistoryUSERDATA**
> GetIsolatedMarginAccountTransferHistoryUSERDATA(ctx, asset, symbol, transFrom, transTo, startTime, endTime, current, size, timestamp, signature, contentType, xMBXAPIKEY)
Get Isolated Margin Account Transfer History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **symbol** | **string**|  | 
  **transFrom** | **string**| \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot; | 
  **transTo** | **string**| \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot; | 
  **startTime** | **string**|  | 
  **endTime** | **string**|  | 
  **current** | **int32**| Currently querying page. Start from 1. Default:1 | 
  **size** | **int32**| Default:10 Max:100 | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
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

# **IsolatedMarginAccountTransferMARGIN**
> IsolatedMarginAccountTransferMARGIN(ctx, asset, symbol, transFrom, transTo, amount, timestamp, signature, contentType, xMBXAPIKEY)
Isolated Margin Account Transfer (MARGIN)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
  **symbol** | **string**| isolated symbol | 
  **transFrom** | **string**| \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot; | 
  **transTo** | **string**| \&quot;SPOT\&quot;, \&quot;ISOLATED_MARGIN\&quot; | 
  **amount** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
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

# **QueryIsolatedMarginAccountInfoUSERDATA**
> QueryIsolatedMarginAccountInfoUSERDATA(ctx, symbols, timestamp, signature, contentType, xMBXAPIKEY)
Query Isolated Margin Account Info (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbols** | **string**| Max 5 symbols can be sent; separated by &#x27;,&#x27;. e.g. &#x27;BTCUSDT,BNBUSDT,ADAUSDT&#x27; | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
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

# **QueryIsolatedMarginPairMARKETDATA**
> QueryIsolatedMarginPairMARKETDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Isolated Margin Pair (MARKET_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**|  | 
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

# **RenewIsolatedMarginAccountListenKey**
> RenewIsolatedMarginAccountListenKey(ctx, listenKey, symbol, contentType, xMBXAPIKEY)
Renew Isolated Margin Account Listen Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listenKey** | **string**|  | 
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

