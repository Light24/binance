# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFiatDepositWithdrawHistoryUSERDATA**](FiatApi.md#GetFiatDepositWithdrawHistoryUSERDATA) | **Get** /sapi/v1/fiat/orders | Get Fiat Deposit/Withdraw History (USER_DATA)
[**GetFiatPaymentsHistoryUSERDATA**](FiatApi.md#GetFiatPaymentsHistoryUSERDATA) | **Get** /sapi/v1/fiat/payments | Get Fiat Payments History (USER_DATA)

# **GetFiatDepositWithdrawHistoryUSERDATA**
> GetFiatDepositWithdrawHistoryUSERDATA(ctx, transactionType, timestamp, signature, contentType, xMBXAPIKEY)
Get Fiat Deposit/Withdraw History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionType** | **int32**| 0-deposit,1-withdraw | 
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

# **GetFiatPaymentsHistoryUSERDATA**
> GetFiatPaymentsHistoryUSERDATA(ctx, transactionType, timestamp, signature, contentType, xMBXAPIKEY)
Get Fiat Payments History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionType** | **int32**| 0-deposit,1-withdraw | 
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

