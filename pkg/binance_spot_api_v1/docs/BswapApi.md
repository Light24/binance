# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLiquidityTRADE**](BswapApi.md#AddLiquidityTRADE) | **Post** /sapi/v1/bswap/liquidityAdd | Add Liquidity (TRADE)
[**GetLiquidityOperationRecordUSERDATA**](BswapApi.md#GetLiquidityOperationRecordUSERDATA) | **Get** /sapi/v1/bswap/liquidityOps | Get Liquidity Operation Record (USER_DATA)
[**GetSwapHistoryUSERDATA**](BswapApi.md#GetSwapHistoryUSERDATA) | **Get** /sapi/v1/bswap/swap | Get Swap History (USER_DATA)
[**GetliquidityinformationofapoolUSERDATA**](BswapApi.md#GetliquidityinformationofapoolUSERDATA) | **Get** /sapi/v1/bswap/liquidity | Get liquidity information of a pool (USER_DATA)
[**ListAllSwapPools**](BswapApi.md#ListAllSwapPools) | **Get** /sapi/v1/bswap/pools | List All Swap Pools
[**RemoveLiquidityTRADE**](BswapApi.md#RemoveLiquidityTRADE) | **Post** /sapi/v1/bswap/liquidityRemove | Remove Liquidity (TRADE)
[**RequestQuoteUSERDATA**](BswapApi.md#RequestQuoteUSERDATA) | **Get** /sapi/v1/bswap/quote | Request Quote (USER_DATA)
[**SwapTRADE**](BswapApi.md#SwapTRADE) | **Post** /sapi/v1/bswap/swap | Swap (TRADE)

# **AddLiquidityTRADE**
> AddLiquidityTRADE(ctx, poolId, asset, quantity, timestamp, signature, contentType, xMBXAPIKEY)
Add Liquidity (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
  **asset** | **string**|  | 
  **quantity** | **string**|  | 
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

# **GetLiquidityOperationRecordUSERDATA**
> GetLiquidityOperationRecordUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Liquidity Operation Record (USER_DATA)

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

# **GetSwapHistoryUSERDATA**
> GetSwapHistoryUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get Swap History (USER_DATA)

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

# **GetliquidityinformationofapoolUSERDATA**
> GetliquidityinformationofapoolUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Get liquidity information of a pool (USER_DATA)

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

# **ListAllSwapPools**
> ListAllSwapPools(ctx, contentType, xMBXAPIKEY)
List All Swap Pools

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

# **RemoveLiquidityTRADE**
> RemoveLiquidityTRADE(ctx, poolId, type_, shareAmount, timestamp, signature, contentType, xMBXAPIKEY)
Remove Liquidity (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
  **type_** | **string**| SINGLE for single asset removal, COMBINATION for combination of all coins removal | 
  **shareAmount** | **string**|  | 
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

# **RequestQuoteUSERDATA**
> RequestQuoteUSERDATA(ctx, quoteAsset, baseAsset, quoteQty, timestamp, signature, contentType, xMBXAPIKEY)
Request Quote (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteAsset** | **string**|  | 
  **baseAsset** | **string**|  | 
  **quoteQty** | **string**|  | 
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

# **SwapTRADE**
> SwapTRADE(ctx, quoteAsset, baseAsset, quoteQty, timestamp, signature, contentType, xMBXAPIKEY)
Swap (TRADE)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteAsset** | **string**|  | 
  **baseAsset** | **string**|  | 
  **quoteQty** | **string**|  | 
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

