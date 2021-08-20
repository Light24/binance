# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeFixedActivityPositiontoDailyPositionUSERDATA**](SavingsApi.md#ChangeFixedActivityPositiontoDailyPositionUSERDATA) | **Post** /sapi/v1/lending/positionChanged | Change Fixed/Activity Position to Daily Position(USER_DATA)
[**GetCustomizedFixedProjectPositionUSERDATA**](SavingsApi.md#GetCustomizedFixedProjectPositionUSERDATA) | **Get** /sapi/v1/lending/project/position/list | Get Customized Fixed Project Position (USER_DATA)
[**GetFixedandCustomizedFixedProjectListUSERDATA**](SavingsApi.md#GetFixedandCustomizedFixedProjectListUSERDATA) | **Get** /sapi/v1/lending/project/list | Get Fixed and Customized Fixed Project List(USER_DATA)
[**GetFlexibleProductListUSERDATA**](SavingsApi.md#GetFlexibleProductListUSERDATA) | **Get** /sapi/v1/lending/daily/product/list | Get Flexible Product List (USER_DATA)
[**GetFlexibleProductPositionUSERDATA**](SavingsApi.md#GetFlexibleProductPositionUSERDATA) | **Get** /sapi/v1/lending/daily/token/position | Get Flexible Product Position (USER_DATA)
[**GetInterestHistoryUSERDATA1**](SavingsApi.md#GetInterestHistoryUSERDATA1) | **Get** /sapi/v1/lending/union/interestHistory | Get Interest History (USER_DATA)
[**GetLeftDailyPurchaseQuotaofFlexibleProductUSERDATA**](SavingsApi.md#GetLeftDailyPurchaseQuotaofFlexibleProductUSERDATA) | **Get** /sapi/v1/lending/daily/userLeftQuota | Get Left Daily Purchase Quota of Flexible Product (USER_DATA)
[**GetLeftDailyRedemptionQuotaofFlexibleProduct**](SavingsApi.md#GetLeftDailyRedemptionQuotaofFlexibleProduct) | **Get** /sapi/v1/lending/daily/userRedemptionQuota | Get Left Daily Redemption Quota of Flexible Product
[**GetPurchaseRecordUSERDATA**](SavingsApi.md#GetPurchaseRecordUSERDATA) | **Get** /sapi/v1/lending/union/purchaseRecord | Get Purchase Record (USER_DATA)
[**GetRedemptionRecordUSERDATA**](SavingsApi.md#GetRedemptionRecordUSERDATA) | **Get** /sapi/v1/lending/union/redemptionRecord | Get Redemption Record (USER_DATA)
[**LendingAccountUSERDATA**](SavingsApi.md#LendingAccountUSERDATA) | **Get** /sapi/v1/lending/union/account | Lending Account (USER_DATA)
[**PurchaseCustomizedFixedProjectUSERDATA**](SavingsApi.md#PurchaseCustomizedFixedProjectUSERDATA) | **Post** /sapi/v1/lending/customizedFixed/purchase | Purchase Customized Fixed Project (USER_DATA)
[**PurchaseFlexibleProductUSERDATA**](SavingsApi.md#PurchaseFlexibleProductUSERDATA) | **Post** /sapi/v1/lending/daily/purchase | Purchase Flexible Product (USER_DATA)
[**RedeemFlexibleProductUSERDATA**](SavingsApi.md#RedeemFlexibleProductUSERDATA) | **Post** /sapi/v1/lending/daily/redeem | Redeem Flexible Product (USER_DATA)

# **ChangeFixedActivityPositiontoDailyPositionUSERDATA**
> ChangeFixedActivityPositiontoDailyPositionUSERDATA(ctx, projectId, lot, timestamp, signature, contentType, xMBXAPIKEY)
Change Fixed/Activity Position to Daily Position(USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **lot** | **string**|  | 
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

# **GetCustomizedFixedProjectPositionUSERDATA**
> GetCustomizedFixedProjectPositionUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Get Customized Fixed Project Position (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
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

# **GetFixedandCustomizedFixedProjectListUSERDATA**
> GetFixedandCustomizedFixedProjectListUSERDATA(ctx, type_, timestamp, signature, contentType, xMBXAPIKEY)
Get Fixed and Customized Fixed Project List(USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| \&quot;ACTIVITY\&quot;, \&quot;CUSTOMIZED_FIXED\&quot; | 
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

# **GetFlexibleProductListUSERDATA**
> GetFlexibleProductListUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Flexible Product List (USER_DATA)

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

# **GetFlexibleProductPositionUSERDATA**
> GetFlexibleProductPositionUSERDATA(ctx, asset, timestamp, signature, contentType, xMBXAPIKEY)
Get Flexible Product Position (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asset** | **string**|  | 
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

# **GetInterestHistoryUSERDATA1**
> GetInterestHistoryUSERDATA1(ctx, lendingType, timestamp, signature, contentType, xMBXAPIKEY)
Get Interest History (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lendingType** | **string**|  | 
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

# **GetLeftDailyPurchaseQuotaofFlexibleProductUSERDATA**
> GetLeftDailyPurchaseQuotaofFlexibleProductUSERDATA(ctx, productId, timestamp, signature, contentType, xMBXAPIKEY)
Get Left Daily Purchase Quota of Flexible Product (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

# **GetLeftDailyRedemptionQuotaofFlexibleProduct**
> GetLeftDailyRedemptionQuotaofFlexibleProduct(ctx, productId, type_, timestamp, signature, contentType, xMBXAPIKEY)
Get Left Daily Redemption Quota of Flexible Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

# **GetPurchaseRecordUSERDATA**
> GetPurchaseRecordUSERDATA(ctx, lendingType, timestamp, signature, contentType, xMBXAPIKEY)
Get Purchase Record (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lendingType** | **string**|  | 
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

# **GetRedemptionRecordUSERDATA**
> GetRedemptionRecordUSERDATA(ctx, lendingType, timestamp, signature, contentType, xMBXAPIKEY)
Get Redemption Record (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lendingType** | **string**|  | 
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

# **LendingAccountUSERDATA**
> LendingAccountUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Lending Account (USER_DATA)

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

# **PurchaseCustomizedFixedProjectUSERDATA**
> PurchaseCustomizedFixedProjectUSERDATA(ctx, projectId, lot, timestamp, signature, contentType, xMBXAPIKEY)
Purchase Customized Fixed Project (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 
  **lot** | **string**|  | 
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

# **PurchaseFlexibleProductUSERDATA**
> PurchaseFlexibleProductUSERDATA(ctx, productId, amount, timestamp, signature, contentType, xMBXAPIKEY)
Purchase Flexible Product (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

# **RedeemFlexibleProductUSERDATA**
> RedeemFlexibleProductUSERDATA(ctx, productId, amount, type_, timestamp, signature, contentType, xMBXAPIKEY)
Redeem Flexible Product (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
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

