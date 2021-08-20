# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountInformationUSERDATA**](TradeApi.md#AccountInformationUSERDATA) | **Get** /api/v3/account | Account Information (USER_DATA)
[**AccountTradeListUSERDATA**](TradeApi.md#AccountTradeListUSERDATA) | **Get** /api/v3/myTrades | Account Trade List (USER_DATA)
[**AllOrdersUSERDATA**](TradeApi.md#AllOrdersUSERDATA) | **Get** /api/v3/allOrders | All Orders (USER_DATA)
[**CancelOCOTRADE**](TradeApi.md#CancelOCOTRADE) | **Delete** /api/v3/orderList | Cancel OCO (TRADE)
[**CancelOrder**](TradeApi.md#CancelOrder) | **Delete** /api/v3/order | Cancel Order
[**CancelallOpenOrdersonaSymbolTRADE**](TradeApi.md#CancelallOpenOrdersonaSymbolTRADE) | **Delete** /api/v3/openOrders | Cancel all Open Orders on a Symbol (TRADE)
[**CurrentOpenOrdersUSERDATA**](TradeApi.md#CurrentOpenOrdersUSERDATA) | **Get** /api/v3/openOrders | Current Open Orders (USER_DATA)
[**NewOCOTRADE**](TradeApi.md#NewOCOTRADE) | **Post** /api/v3/order/oco | New OCO (TRADE)
[**NewOrder**](TradeApi.md#NewOrder) | **Post** /api/v3/order | New Order
[**QueryOCOUSERDATA**](TradeApi.md#QueryOCOUSERDATA) | **Get** /api/v3/orderList | Query OCO (USER_DATA)
[**QueryOpenOCOUSERDATA**](TradeApi.md#QueryOpenOCOUSERDATA) | **Get** /api/v3/openOrderList | Query Open OCO (USER_DATA)
[**QueryOrderUSERDATA**](TradeApi.md#QueryOrderUSERDATA) | **Get** /api/v3/order | Query Order (USER_DATA)
[**QueryallOCOUSERDATA**](TradeApi.md#QueryallOCOUSERDATA) | **Get** /api/v3/allOrderList | Query all OCO (USER_DATA)
[**TestNewOrderTRADE**](TradeApi.md#TestNewOrderTRADE) | **Post** /api/v3/order/test | Test New Order (TRADE)

# **AccountInformationUSERDATA**
> AccountInformationUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Account Information (USER_DATA)

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

# **AccountTradeListUSERDATA**
> AccountTradeListUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Account Trade List (USER_DATA)

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

# **AllOrdersUSERDATA**
> AllOrdersUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
All Orders (USER_DATA)

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

# **CancelOCOTRADE**
> CancelOCOTRADE(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Cancel OCO (TRADE)

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

# **CancelOrder**
> CancelOrder(ctx, symbol, orderId, timestamp, signature, contentType, xMBXAPIKEY)
Cancel Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **orderId** | **string**|  | 
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

# **CancelallOpenOrdersonaSymbolTRADE**
> CancelallOpenOrdersonaSymbolTRADE(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Cancel all Open Orders on a Symbol (TRADE)

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

# **CurrentOpenOrdersUSERDATA**
> CurrentOpenOrdersUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Current Open Orders (USER_DATA)

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

# **NewOCOTRADE**
> NewOCOTRADE(ctx, symbol, side, quantity, price, stopPrice, timestamp, signature, contentType, xMBXAPIKEY)
New OCO (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **side** | **string**|  | 
  **quantity** | **string**|  | 
  **price** | **string**|  | 
  **stopPrice** | **string**|  | 
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

# **NewOrder**
> NewOrder(ctx, symbol, side, type_, quantity, price, timestamp, signature, contentType, xMBXAPIKEY)
New Order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **side** | **string**|  | 
  **type_** | **string**|  | 
  **quantity** | **float64**|  | 
  **price** | **int32**|  | 
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

# **QueryOCOUSERDATA**
> QueryOCOUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query OCO (USER_DATA)

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

# **QueryOpenOCOUSERDATA**
> QueryOpenOCOUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Open OCO (USER_DATA)

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

# **QueryOrderUSERDATA**
> QueryOrderUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Order (USER_DATA)

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

# **QueryallOCOUSERDATA**
> QueryallOCOUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query all OCO (USER_DATA)

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

# **TestNewOrderTRADE**
> TestNewOrderTRADE(ctx, symbol, side, type_, quantity, price, timestamp, signature, contentType, xMBXAPIKEY)
Test New Order (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **side** | **string**|  | 
  **type_** | **string**|  | 
  **quantity** | **float64**|  | 
  **price** | **int32**|  | 
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

