# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetC2CTradeHistoryUSERDATA**](C2CApi.md#GetC2CTradeHistoryUSERDATA) | **Get** /sapi/v1/c2c/orderMatch/listUserOrderHistory | Get C2C Trade History (USER_DATA)

# **GetC2CTradeHistoryUSERDATA**
> GetC2CTradeHistoryUSERDATA(ctx, tradeType, timestamp, signature, contentType, xMBXAPIKEY)
Get C2C Trade History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeType** | **string**| BUY, SELL | 
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

