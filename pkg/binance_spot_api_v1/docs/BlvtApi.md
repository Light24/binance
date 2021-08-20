# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBLVTUserLimitInfo**](BlvtApi.md#GetBLVTUserLimitInfo) | **Get** /sapi/v1/blvt/userLimit | Get BLVT User Limit Info
[**Leveragetokeninfo**](BlvtApi.md#Leveragetokeninfo) | **Get** /sapi/v1/blvt/tokenInfo | Leverage token info
[**QueryRedemptionRecord**](BlvtApi.md#QueryRedemptionRecord) | **Get** /sapi/v1/blvt/redeem/record | Query Redemption Record
[**QuerySubscriptionRecord**](BlvtApi.md#QuerySubscriptionRecord) | **Get** /sapi/v1/blvt/subscribe/record | Query Subscription Record
[**RedeemBLVT**](BlvtApi.md#RedeemBLVT) | **Post** /sapi/v1/blvt/redeem | Redeem BLVT
[**SubscribeBLVT**](BlvtApi.md#SubscribeBLVT) | **Post** /sapi/v1/blvt/subscribe | Subscribe BLVT

# **GetBLVTUserLimitInfo**
> GetBLVTUserLimitInfo(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get BLVT User Limit Info

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

# **Leveragetokeninfo**
> Leveragetokeninfo(ctx, contentType, xMBXAPIKEY)
Leverage token info

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

# **QueryRedemptionRecord**
> QueryRedemptionRecord(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Redemption Record

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

# **QuerySubscriptionRecord**
> QuerySubscriptionRecord(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Query Subscription Record

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

# **RedeemBLVT**
> RedeemBLVT(ctx, tokenName, amount, timestamp, signature, contentType, xMBXAPIKEY)
Redeem BLVT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenName** | **string**|  | 
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

# **SubscribeBLVT**
> SubscribeBLVT(ctx, tokenName, cost, timestamp, signature, contentType, xMBXAPIKEY)
Subscribe BLVT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenName** | **string**|  | 
  **cost** | **string**|  | 
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

