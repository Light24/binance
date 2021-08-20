# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountTradeList**](TradeApi.md#AccountTradeList) | **Get** /dapi/v1/userTrades | Account Trade List
[**Accountinformation**](TradeApi.md#Accountinformation) | **Get** /dapi/v1/account | Account information
[**AllOrders**](TradeApi.md#AllOrders) | **Get** /dapi/v1/allOrders | All Orders
[**AutoCancelAllOpenOrdersTRADE**](TradeApi.md#AutoCancelAllOpenOrdersTRADE) | **Post** /dapi/v1/countdownCancelAll | Auto-Cancel All Open Orders (TRADE)
[**CancelAllOpenOrders**](TradeApi.md#CancelAllOpenOrders) | **Delete** /dapi/v1/allOpenOrders | Cancel All Open Orders
[**CancelMultipleOrdersTRADE**](TradeApi.md#CancelMultipleOrdersTRADE) | **Delete** /dapi/v1/batchOrders | Cancel Multiple Orders (TRADE)
[**CancelOrder**](TradeApi.md#CancelOrder) | **Delete** /dapi/v1/order | Cancel Order
[**ChangeInitialLeverage**](TradeApi.md#ChangeInitialLeverage) | **Post** /dapi/v1/leverage | Change Initial Leverage
[**ChangeMarginType**](TradeApi.md#ChangeMarginType) | **Post** /dapi/v1/marginType | Change Margin Type
[**ChangePositionModeTRADE**](TradeApi.md#ChangePositionModeTRADE) | **Post** /dapi/v1/positionSide/dual | Change Position Mode（TRADE）
[**CurrentAllOpenOrdersUSERDATA**](TradeApi.md#CurrentAllOpenOrdersUSERDATA) | **Get** /dapi/v1/openOrders | Current All Open Orders (USER_DATA)
[**GetCurrentPositionModeUSERDATA**](TradeApi.md#GetCurrentPositionModeUSERDATA) | **Get** /dapi/v1/positionSide/dual | Get Current Position Mode（USER_DATA）
[**GetIncomeHistory**](TradeApi.md#GetIncomeHistory) | **Get** /dapi/v1/income | Get Income History
[**GetPostionMarginChangeHistory**](TradeApi.md#GetPostionMarginChangeHistory) | **Get** /dapi/v1/positionMargin/history | Get Postion Margin Change History
[**ModifyIsolatedPositionMargin**](TradeApi.md#ModifyIsolatedPositionMargin) | **Post** /dapi/v1/positionMargin | Modify Isolated Position Margin
[**NewOrderTRADE**](TradeApi.md#NewOrderTRADE) | **Post** /dapi/v1/order | New Order (TRADE)
[**NotionalBracketUSERDATA**](TradeApi.md#NotionalBracketUSERDATA) | **Get** /dapi/v1/leverageBracket | Notional Bracket (USER_DATA)
[**PlaceMultipleOrdersTRADE**](TradeApi.md#PlaceMultipleOrdersTRADE) | **Post** /dapi/v1/batchOrders | Place Multiple Orders (TRADE)
[**PositionADLQuantileEstimationUSERDATA**](TradeApi.md#PositionADLQuantileEstimationUSERDATA) | **Get** /dapi/v1/adlQuantile | Position ADL Quantile Estimation (USER_DATA)
[**PositionInformation**](TradeApi.md#PositionInformation) | **Get** /dapi/v1/positionRisk | Position Information
[**QueryCurrentOpenOrderUSERDATA**](TradeApi.md#QueryCurrentOpenOrderUSERDATA) | **Get** /dapi/v1/openOrder | Query Current Open Order (USER_DATA)
[**QueryOrder**](TradeApi.md#QueryOrder) | **Get** /dapi/v1/order | Query Order
[**UsersForceOrdersUSERDATA**](TradeApi.md#UsersForceOrdersUSERDATA) | **Get** /dapi/v1/forceOrders | User&#x27;s Force Orders (USER_DATA)

# **AccountTradeList**
> AccountTradeList(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Account Trade List

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

# **Accountinformation**
> Accountinformation(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Account information

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

# **AllOrders**
> AllOrders(ctx, timestamp, signature, contentType, xMBXAPIKEY)
All Orders

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
> CancelMultipleOrdersTRADE(ctx, symbol, orderIdList, timestamp, signature, contentType, xMBXAPIKEY)
Cancel Multiple Orders (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **orderIdList** | **string**| [id,id]  no space in between | 
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

# **ChangeInitialLeverage**
> ChangeInitialLeverage(ctx, symbol, leverage, timestamp, signature, contentType, xMBXAPIKEY)
Change Initial Leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **leverage** | **string**|  | 
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

# **ChangeMarginType**
> ChangeMarginType(ctx, symbol, marginType, timestamp, signature, contentType, xMBXAPIKEY)
Change Margin Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **marginType** | **string**|  | 
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

# **ChangePositionModeTRADE**
> ChangePositionModeTRADE(ctx, dualSidePosition, timestamp, signature, contentType, xMBXAPIKEY)
Change Position Mode（TRADE）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dualSidePosition** | **string**| Boolean value. \&quot;true\&quot;: Hedge Mode mode; \&quot;false\&quot;: One-way Mode | 
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

# **GetCurrentPositionModeUSERDATA**
> GetCurrentPositionModeUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Current Position Mode（USER_DATA）

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

# **GetIncomeHistory**
> GetIncomeHistory(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Income History

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

# **GetPostionMarginChangeHistory**
> GetPostionMarginChangeHistory(ctx, symbol, timestamp, signature, contentType, xMBXAPIKEY)
Get Postion Margin Change History

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

# **ModifyIsolatedPositionMargin**
> ModifyIsolatedPositionMargin(ctx, symbol, amount, type_, timestamp, signature, contentType, xMBXAPIKEY)
Modify Isolated Position Margin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **amount** | **string**|  | 
  **type_** | **string**|  | 
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

# **NotionalBracketUSERDATA**
> NotionalBracketUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Notional Bracket (USER_DATA)

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

# **PositionADLQuantileEstimationUSERDATA**
> PositionADLQuantileEstimationUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Position ADL Quantile Estimation (USER_DATA)

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

# **PositionInformation**
> PositionInformation(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Position Information

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

# **QueryCurrentOpenOrderUSERDATA**
> QueryCurrentOpenOrderUSERDATA(ctx, symbol, orderId, timestamp, signature, contentType, xMBXAPIKEY)
Query Current Open Order (USER_DATA)

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

# **QueryOrder**
> QueryOrder(ctx, symbol, orderId, timestamp, signature, contentType, xMBXAPIKEY)
Query Order

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

