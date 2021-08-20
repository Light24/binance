# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeInitialLeverage**](TradeApi.md#ChangeInitialLeverage) | **Post** /fapi/v1/leverage | Change Initial Leverage
[**ChangeMarginType**](TradeApi.md#ChangeMarginType) | **Post** /fapi/v1/marginType | Change Margin Type
[**ChangePositionModeTRADE**](TradeApi.md#ChangePositionModeTRADE) | **Post** /fapi/v1/positionSide/dual | Change Position Mode（TRADE）
[**GetCurrentPositionModeUSERDATA**](TradeApi.md#GetCurrentPositionModeUSERDATA) | **Get** /fapi/v1/positionSide/dual | Get Current Position Mode（USER_DATA）
[**GetPostionMarginChangeHistory**](TradeApi.md#GetPostionMarginChangeHistory) | **Get** /fapi/v1/positionMargin/history | Get Postion Margin Change History
[**ModifyIsolatedPositionMargin**](TradeApi.md#ModifyIsolatedPositionMargin) | **Post** /fapi/v1/positionMargin | Modify Isolated Position Margin
[**UserAPITradingQuantitativeRulesIndicatorsUSERDATA**](TradeApi.md#UserAPITradingQuantitativeRulesIndicatorsUSERDATA) | **Get** /fapi/v1/apiTradingStatus | User API Trading Quantitative Rules Indicators (USER_DATA)

# **ChangeInitialLeverage**
> ChangeInitialLeverage(ctx, symbol, leverage, timestamp, signature, contentType, xMBXAPIKEY)
Change Initial Leverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **symbol** | **string**|  | 
  **leverage** | **int32**|  | 
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
  **amount** | **int32**|  | 
  **type_** | **int32**|  | 
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

# **UserAPITradingQuantitativeRulesIndicatorsUSERDATA**
> UserAPITradingQuantitativeRulesIndicatorsUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
User API Trading Quantitative Rules Indicators (USER_DATA)

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

