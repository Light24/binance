# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountListUSERDATA**](MiningApi.md#AccountListUSERDATA) | **Get** /sapi/v1/mining/statistics/user/list | Account List (USER_DATA)
[**AcquiringAlgorithmMARKETDATA**](MiningApi.md#AcquiringAlgorithmMARKETDATA) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm (MARKET_DATA)
[**AcquiringCoinNameMARKETDATA**](MiningApi.md#AcquiringCoinNameMARKETDATA) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName (MARKET_DATA)
[**CancelHashrateResaleconfiguration**](MiningApi.md#CancelHashrateResaleconfiguration) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel Hashrate Resale configuration
[**EarningsListUSERDATA**](MiningApi.md#EarningsListUSERDATA) | **Get** /sapi/v1/mining/payment/list | Earnings List (USER_DATA)
[**ExtraBonusListUSERDATA**](MiningApi.md#ExtraBonusListUSERDATA) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List (USER_DATA)
[**HashrateResaleDetailsUSERDATA**](MiningApi.md#HashrateResaleDetailsUSERDATA) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Details (USER_DATA)
[**HashrateResaleListUSERDATA**](MiningApi.md#HashrateResaleListUSERDATA) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List (USER_DATA)
[**HashrateResaleRequestUSERDATA**](MiningApi.md#HashrateResaleRequestUSERDATA) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request (USER_DATA)
[**RequestforDetailMinerListUSERDATA**](MiningApi.md#RequestforDetailMinerListUSERDATA) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List (USER_DATA)
[**RequestforMinerListUSERDATA**](MiningApi.md#RequestforMinerListUSERDATA) | **Get** /sapi/v1/mining/worker/list | Request for Miner List (USER_DATA)
[**StatisticListUSERDATA**](MiningApi.md#StatisticListUSERDATA) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List (USER_DATA)

# **AccountListUSERDATA**
> AccountListUSERDATA(ctx, algo, userName, timestamp, signature, contentType, xMBXAPIKEY)
Account List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
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

# **AcquiringAlgorithmMARKETDATA**
> AcquiringAlgorithmMARKETDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Acquiring Algorithm (MARKET_DATA)

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

# **AcquiringCoinNameMARKETDATA**
> AcquiringCoinNameMARKETDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Acquiring CoinName (MARKET_DATA)

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

# **CancelHashrateResaleconfiguration**
> CancelHashrateResaleconfiguration(ctx, configId, userName, timestamp, signature, contentType, xMBXAPIKEY)
Cancel Hashrate Resale configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 
  **userName** | **string**|  | 
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

# **EarningsListUSERDATA**
> EarningsListUSERDATA(ctx, algo, userName, timestamp, signature, contentType, xMBXAPIKEY)
Earnings List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
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

# **ExtraBonusListUSERDATA**
> ExtraBonusListUSERDATA(ctx, algo, userName, timestamp, signature, contentType, xMBXAPIKEY)
Extra Bonus List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
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

# **HashrateResaleDetailsUSERDATA**
> HashrateResaleDetailsUSERDATA(ctx, configId, userName, timestamp, signature, contentType, xMBXAPIKEY)
Hashrate Resale Details (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 
  **userName** | **string**|  | 
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

# **HashrateResaleListUSERDATA**
> HashrateResaleListUSERDATA(ctx, timestamp, signature, contentType, xMBXAPIKEY)
Hashrate Resale List (USER_DATA)

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

# **HashrateResaleRequestUSERDATA**
> HashrateResaleRequestUSERDATA(ctx, userName, algo, startDate, endDate, toPoolUser, hashRate, timestamp, signature, contentType, xMBXAPIKEY)
Hashrate Resale Request (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**|  | 
  **algo** | **string**|  | 
  **startDate** | **string**|  | 
  **endDate** | **string**|  | 
  **toPoolUser** | **string**|  | 
  **hashRate** | **string**|  | 
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

# **RequestforDetailMinerListUSERDATA**
> RequestforDetailMinerListUSERDATA(ctx, algo, userName, workerName, timestamp, signature, contentType, xMBXAPIKEY)
Request for Detail Miner List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
  **workerName** | **string**| Miner’s name（required） | 
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

# **RequestforMinerListUSERDATA**
> RequestforMinerListUSERDATA(ctx, algo, userName, timestamp, signature, contentType, xMBXAPIKEY)
Request for Miner List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
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

# **StatisticListUSERDATA**
> StatisticListUSERDATA(ctx, algo, userName, timestamp, signature, contentType, xMBXAPIKEY)
Statistic List (USER_DATA)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **algo** | **string**| Algorithm(sha256) | 
  **userName** | **string**| Mining account | 
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

