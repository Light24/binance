# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllOrders**](OrderApi.md#AllOrders) | **Get** /fapi/v1/allOrders | All Orders
[**AutoCancelAllOpenOrdersTRADE**](OrderApi.md#AutoCancelAllOpenOrdersTRADE) | **Post** /fapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**CancelAllOpenOrders**](OrderApi.md#CancelAllOpenOrders) | **Delete** /fapi/v1/allOpenOrders | Cancel All Open Orders
[**CancelMultipleOrdersTRADE**](OrderApi.md#CancelMultipleOrdersTRADE) | **Delete** /fapi/v1/batchOrders | Cancel Multiple Orders (TRADE)
[**CancelOrder**](OrderApi.md#CancelOrder) | **Delete** /fapi/v1/order | Cancel Order
[**CurrentAllOpenOrdersUSERDATA**](OrderApi.md#CurrentAllOpenOrdersUSERDATA) | **Get** /fapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**NewOrderTRADE**](OrderApi.md#NewOrderTRADE) | **Post** /fapi/v1/order | New Order (TRADE)
[**PlaceMultipleOrdersTRADE**](OrderApi.md#PlaceMultipleOrdersTRADE) | **Post** /fapi/v1/batchOrders | Place Multiple Orders (TRADE)
[**QueryCurrentOpenOrderUSERDATA**](OrderApi.md#QueryCurrentOpenOrderUSERDATA) | **Get** /fapi/v1/openOrder | Query Current Open Order (USER_DATA)
[**QueryOrder**](OrderApi.md#QueryOrder) | **Get** /fapi/v1/order | Query Order
[**UsersForceOrdersUSERDATA**](OrderApi.md#UsersForceOrdersUSERDATA) | **Get** /fapi/v1/forceOrders | User&#x27;s Force Orders (USER_DATA)

# **AllOrders**
> AllOrders(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
All Orders

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

# **AutoCancelAllOpenOrdersTRADE**
> AutoCancelAllOpenOrdersTRADE(ctx, symbol, countdownTime, timestamp, signature, contentType, xMBXAPIKEY)
Auto-Cancel All Open Orders (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **countdownTime** | **string**| countdown time, 1000 for 1 second. 0 to cancel the timer | 
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

# **CancelAllOpenOrders**
> CancelAllOpenOrders(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Cancel All Open Orders

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

# **CancelMultipleOrdersTRADE**
> CancelMultipleOrdersTRADE(ctx, symbol, origClientOrderIdList, timestamp, signature, contentType, xMBXAPIKEY)
Cancel Multiple Orders (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **origClientOrderIdList** | **string**|  | 
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
> CancelOrder(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Cancel Order

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

# **CurrentAllOpenOrdersUSERDATA**
> CurrentAllOpenOrdersUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Current All Open Orders (USER_DATA)

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

# **NewOrderTRADE**
> NewOrderTRADE(ctx, symbol, side, type_, quantity, price, timestamp, signature, contentType, xMBXAPIKEY)
New Order (TRADE)

https://binance-docs.github.io/apidocs/futures/en/#new-order-trade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **side** | **string**|  | 
  **type_** | **string**|  | 
  **quantity** | **string**|  | 
  **price** | **string**|  | 
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

# **PlaceMultipleOrdersTRADE**
> PlaceMultipleOrdersTRADE(ctx, batchOrders, timestamp, signature, contentType, xMBXAPIKEY)
Place Multiple Orders (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchOrders** | **string**|  | 
  **timestamp** | **string**|  | 
  **signature** | **string**| Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent with Hedge Mode. | 
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

# **QueryCurrentOpenOrderUSERDATA**
> QueryCurrentOpenOrderUSERDATA(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Current Open Order (USER_DATA)

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

# **QueryOrder**
> QueryOrder(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Query Order

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

# **UsersForceOrdersUSERDATA**
> UsersForceOrdersUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
User's Force Orders (USER_DATA)

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

